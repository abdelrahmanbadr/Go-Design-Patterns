Singleton
=============

**THIS IS CONSIDERED TO BE AN ANTI-PATTERN! FOR BETTER TESTABILITY AND
MAINTAINABILITY USE DEPENDENCY INJECTION!**

Purpose
-------

To have only one instance of this object in the application that will
handle all calls.

Examples
--------

-  DB Connector
-  Game board in a monopoly game
-  Maze in a maze-game
-  Logger
-  Lock file for the application (there is only one in the filesystem
   ...)

Test
----

singletone_test.go