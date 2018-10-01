`Prototype`__
=============

Purpose
-------

To avoid the cost of creating objects the standard way and
instead create a prototype and clone it.

This pattern is used when creation of object directly is costly.
For example, an object is to be created after a costly database operation.

Examples
--------

-  Large amounts of data (e.g. create 1,000,000 rows in a database at
   once via a ORM).

Test
----

prototype_test.php
