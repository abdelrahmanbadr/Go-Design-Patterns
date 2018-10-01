Builder
===========

Purpose
-------

Builder is an interface that build parts of a complex object.

Sometimes, if the builder has a better knowledge of what it builds, this
interface could be an abstract class with default methods (aka adapter).

If you have a complex inheritance tree for objects, it is logical to
have a complex inheritance tree for builders too.


UML Diagram
-----------

.. image:: uml/builder.png
   :alt: Alt Builder UML Diagram
   :align: center

Test
----

builder_test.go