# DEDIS framework splitted in modules

The purpose of this repository is to give a hint of how the Cothority framework
could be splitted into smaller modules that would replacable more easily because
the current situation is that changing part of the framework is nearly
impossible without changes from top to bottom.

It would be help when planning student projects (PhD students too) as they could
focus on one or a few modules. It would also help the creation of a unit test
suite that could replace the current Travis workflow to bring it down to a few
minutes duration instead of an hour.

## The Big Picture

Each module provides a public API that other modules can use and that's it, no
access to the internal logic.

For instance for the overlay network module, instead of having to declare
services and then protocols, the module provides the creation of RPCs in
multiple namespaces so that they interact with their pair in each node in
the sense that a given RPC will receive messages from the same instance
replicated on each node, and sending messages the same way. No intermodule
communication. Then each RPC can be call either on multiple peers or on a
single one. Think about CoSi. The overlay takes care of gathering the replies
in an optimize way to every online node.

The skipchain module is essentially just a secure distributed storage thus the
only actions needed are a store and get (with proof) functions. Of course the
implementation of the interface will have some requirements like a validation
function but this is not related to the API, but to the instantiation.

The Byzcoin module needs to provide a function to add a transaction to the chain
and a way to get past/current states with the proof and observe what is happening
on the chain.

Internally, those modules are using other modules like a global state storage
module that can be used to either update or read the global state. The
implementation can decide how it is stored. Another module is needed for the
access control (think about DARCs).

The repository tries to enlight those aspects by showing that it is possible
to split the framework with a bit of work.
