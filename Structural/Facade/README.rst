Facade
==========

Purpose
-------
Facade pattern hides the complexities of the system.

The primary goal of a Facade Pattern is not to avoid you having to read the
manual of a complex API. It's only a side-effect. The first goal is to
reduce coupling and follow the Law of Demeter.

A Facade is meant to decouple a client and a sub-system by embedding
many (but sometimes just one) interface, and of course to reduce
complexity.

-  A facade does not forbid you the access to the sub-system
-  You can (you should) have multiple facades for one sub-system


UML Diagram
-----------

.. image:: uml/facade.jpg
   :alt: Alt Facade UML Diagram
   :align: center

Test
----
facade_test.go
