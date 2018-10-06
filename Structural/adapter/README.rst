Adapter / Wrapper
=====================
Introduction
-------
The Adapter Pattern is responsible for adaptation of two incompatible interfaces. It is a structural pattern that is responsible to join functionalities of independent or incompatible interfaces without modifing their implementation.

Interfaces may be incompatible but the inner functionality should suit the need. It allows otherwise incompatible objects to work together by converting the interface of each struct into an interface expected by the clients.

Purpose
-------
- Adapter lets objects work together, that could not otherwise because of incompatible interfaces.
- Wrap the interface of a object into another interface clients expect.
- Allows classes to work together that normally could not because of incompatible interfaces by providing its interface to clients while using the original interface.

Examples
--------

-  DB Client libraries adapter
-  using multiple different webservices and adapters normalize data so
   that the outcome is the same for all

UML Diagram
-----------

.. image:: uml/adapter.png
   :alt: Alt Adapter UML Diagram
   :align: center

Tests/AdapterTest.php

adapter_test.go
