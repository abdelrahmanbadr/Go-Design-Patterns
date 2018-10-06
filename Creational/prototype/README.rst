Prototype
=============
The Prototype Pattern creates duplicate objects while keeping performance in mind.
For instance, an object is to be created after a costly database operation. We can cache the object, returns its clone on next request and update the database as and when needed thus reducing database calls.

Purpose
-------

- To avoid the cost of creating objects the standard way and instead create a prototype and clone it.

This pattern is used when creation of object directly is costly.
For example, an object is to be created after a costly database operation.

Terminology
-----------
- Prototype declares an interface for cloning itself
- ConcretePrototype implements an operation for cloning itself
- Client creates a new object by asking a prototype to clone itself

Verdict
-------
One of the disadvantages of this pattern is that the process of copying an object can be complicated.
In addition, structs that have circular references to other classes are difficult to clone.
Its overuse could affect performance, as the prototype object itself would need to be instantiated if you use a registry of prototypes.

In the context of Golang, I don’t see any particular reason to adopt it.
Golang already provides builtin mechanism for cloning objects.

Examples
--------

-  Large amounts of data (e.g. create 1,000,000 rows in a database at once via a ORM).

Test
----

prototype_test.php
