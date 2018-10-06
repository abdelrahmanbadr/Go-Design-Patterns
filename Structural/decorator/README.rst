Decorator
=============
The Decorator pattern adds new functionality to an existing object without altering its structure.

Purpose
-------
- Dynamically add new functionality to class instances.
- Decorators provide a flexible alternative to inheritance for extending functionality.
- Wrapping a present, putting it in a box, and wrapping the box.

Terminology
-----------
- "Component" defines the interface for objects that can have responsibilities added to them dynamically.

Examples
--------

-  Web Service Layer: Decorators JSON and XML for a REST service (in this case, only one of these should be allowed of course)

UML Diagram
-----------

.. image:: uml/decorator.png
   :alt: Alt Decorator UML Diagram
   :align: center


Test
----

decorator_test.go
