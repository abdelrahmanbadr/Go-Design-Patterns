Adapter
-------
Adapts a given class/object to a new interface. The problem we are solving here is that of non-compatible interfaces.

Facade
-------
Is more like a simple gateway to a complicated set of functionality. You make a black-box for your clients to worry less i.e. make interfaces simpler. 
(make library)
-------

Proxy
-------
Provides the same interface as the proxied-for class. You are solving the problem of the client from having to manage a heavy and/or complex object. (create an object wrapper to cover the main object’s complexity from the client)

Decorator
-------
Is used to add more gunpowder to your objects (note the term objects -- you typically decorate objects dynamically at runtime). You do not hide/impair the existing interfaces of the object but simply extend it at runtime
