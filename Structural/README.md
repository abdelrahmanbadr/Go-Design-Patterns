Adapter
-------
adapts a given class/object to a new interface. In the case of the former, multiple inheritance is typically employed. In the latter case, the object is wrapped by a conforming adapter object and passed around. The problem we are solving here is that of non-compatible interfaces.
Facade
-------
is more like a simple gateway to a complicated set of functionality. You make a black-box for your clients to worry less i.e. make interfaces simpler.
Proxy
-------
provides the same interface as the proxied-for class and typically does some housekeeping stuff on its own. (So instead of making multiple copies of a heavy object X you make copies of a lightweight proxy P which in turn manages X and translates your calls as required.) You are solving the problem of the client from having to manage a heavy and/or complex object.
Decorator
-------
is used to add more gunpowder to your objects (note the term objects -- you typically decorate objects dynamically at runtime). You do not hide/impair the existing interfaces of the object but simply extend it at runtime
