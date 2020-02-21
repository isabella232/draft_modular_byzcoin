package pedersen

import (
	"context"
	"errors"
	"log"
	"time"

	"go.dedis.ch/phoenix/onet"
	"go.dedis.ch/phoenix/types"

	"go.dedis.ch/kyber/v3"
	dkgpedersen "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	vss "go.dedis.ch/kyber/v3/share/vss/pedersen"
	"go.dedis.ch/kyber/v3/suites"
	"go.dedis.ch/kyber/v3/util/key"
)

// Suite is the Kyber suite for Pedersen.
var Suite = suites.MustFind("Ed25519")

type handler struct {
	kp         *key.Pair
	publicKeys []kyber.Point
	onet.DefaultHandler
}

func newHandler(kp *key.Pair, publicKeys []kyber.Point) *handler {
	return &handler{
		kp:         kp,
		publicKeys: publicKeys,
	}
}

func (h handler) Stream(sender onet.Sender, receiver onet.Receiver) error {
	from, req, err := receiver.Recv(context.Background())
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	init, ok := req.(*types.DKGInit)
	if !ok {
		return errors.New("expect init message")
	}

	log.Println("DKG Pedersen init")

	dkg, err := dkgpedersen.NewDistKeyGenerator(Suite, h.kp.Private, h.publicKeys, len(h.publicKeys))
	if err != nil {
		return err
	}

	proc := &processor{
		dkg:    dkg,
		sender: sender,
	}

	err = proc.sendDeals(init)
	if err != nil {
		return err
	}

	for !dkg.Certified() {
		_, req, err := receiver.Recv(ctx)
		if err != nil {
			return err
		}

		switch msg := req.(type) {
		case *types.DKGDeal:
			err = proc.processDeal(init, msg)
		case *types.DKGResponse:
			err = proc.processResponse(msg)
		}

		if err != nil {
			log.Printf("Error during DKG: %+v", err)
		}
	}

	share, err := dkg.DistKeyShare()
	if err != nil {
		return err
	}

	buffer, err := share.Public().MarshalBinary()
	if err != nil {
		return err
	}

	sender.Send(&types.DKGDone{PublicKey: buffer}, from)

	return nil
}

type processor struct {
	dkg    *dkgpedersen.DistKeyGenerator
	sender onet.Sender
}

func (p *processor) sendDeals(init *types.DKGInit) error {
	deals, err := p.dkg.Deals()
	if err != nil {
		return err
	}

	for i, deal := range deals {
		msg := &types.DKGDeal{
			Index: deal.Index,
			Deal: &types.EncryptedDeal{
				DHKey:     deal.Deal.DHKey,
				Signature: deal.Deal.Signature,
				Nonce:     deal.Deal.Nonce,
				Cipher:    deal.Deal.Cipher,
			},
			Signature: deal.Signature,
		}

		err = p.sender.Send(msg, init.GetAddresses()[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *processor) processDeal(init *types.DKGInit, msg *types.DKGDeal) error {
	deal := &dkgpedersen.Deal{
		Index: msg.GetIndex(),
		Deal: &vss.EncryptedDeal{
			DHKey:     msg.GetDeal().GetDHKey(),
			Signature: msg.GetDeal().GetSignature(),
			Nonce:     msg.GetDeal().GetNonce(),
			Cipher:    msg.GetDeal().GetCipher(),
		},
		Signature: msg.GetSignature(),
	}

	resp, err := p.dkg.ProcessDeal(deal)
	if err != nil {
		return err
	}

	respm := &types.DKGResponse{
		Index: resp.Index,
		Response: &types.DKGResponse_Response{
			SessionID: resp.Response.SessionID,
			Index:     resp.Response.Index,
			Signature: resp.Response.Signature,
			Status:    resp.Response.Status,
		},
	}

	err = p.sender.Send(respm, init.GetAddresses()...)
	if err != nil {
		return err
	}

	return nil
}

func (p *processor) processResponse(msg *types.DKGResponse) error {
	resp := &dkgpedersen.Response{
		Index: msg.Index,
		Response: &vss.Response{
			SessionID: msg.GetResponse().GetSessionID(),
			Index:     msg.GetResponse().GetIndex(),
			Signature: msg.GetResponse().GetSignature(),
			Status:    msg.GetResponse().GetStatus(),
		},
	}

	justif, err := p.dkg.ProcessResponse(resp)
	if err != nil {
		return err
	}

	if justif != nil {
		return errors.New("got a justification")
	}

	return nil
}
