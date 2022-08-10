HTTP Server Generator
&middot;
[![Latest Github release](https://img.shields.io/github/release/Martinsos/edlib.svg)](https://github.com/Martinsos/edlib/releases/latest)
[![Build status of the master branch on Linux/OSX](https://img.shields.io/travis/Martinsos/edlib/master?label=Linux%20%2F%20OSX%20build)](https://travis-ci.com/Martinsos/edlib)
[![Build status of the master branch on Windows](https://img.shields.io/appveyor/build/Martinsos/edlib/master?label=Windows%20build)](https://ci.appveyor.com/project/Martinsos/edlib/branch/master)
[![Chat on Gitter](https://img.shields.io/gitter/room/Martinsos/edlib.svg?colorB=753a88)](https://gitter.im/Martinsos/edlib)
[![Published in Bioinformatics](https://img.shields.io/badge/Published%20in-Bioinformatics-167DA4.svg)](https://doi.org/10.1093/bioinformatics/btw753)
=====

A lightweight and super fast Go tool for generating a simple http server. Mainly meant for use in prototyping and hackathons

Creating a dart and go HTTP server is as simple as:
```c
<INSERT COMMAND>
```

## Features
* None.

## Future Features
* Passing in of a csv file to generate both dart reciever and go server
* Creation of custom dart and go models with 1 to 1 comparison 
* Comments mirrored in each codebase 
* Automatic restriction on illegal call types 

## Potential Future Features
* CLI tool for ease of use 
* Automatic docker file creation
* Automatic Swagger support and generation
* Automatic DB setup (most liekly PSQL and MongoDB support)
* Automatic status code correction
* Automatic test file creation
* Interpretation of an existing go server and creating a dart reciever from this 
* Support default http and dio mode in dart (through run flag)
* GraphQL support 
* GRPC support 
* Javascript/typescript server support

## Input
    Service name 
    Get Post etc…
    Query params (key value)
    Query body (key value or json ) [maybe just json]

## Output
  Dart (meant to be inserted into an existing flutter app)
    All functions for each call seperted into files of each call type (GET, PUT, POST, etc...)
    A folder with models
    An empty var with server name (global to services)
  Go 
    Functionaly complete http server
    .env file 
    go.mod 
    go.sum
    File structure and hierachy

## Restrictions
Gets has no body 

### Approach #1: Directly copying edlib source and header files.
Here we directly copied [edlib/](edlib/) directory to our project, to get following project structure:
```
edlib/  -> copied from edlib/
  include/
    edlib.h
  src/
    edlib.cpp
helloWorld.cpp -> your program
```

Since `helloWorld` is a c++ program, we can compile it with just one line: `c++ helloWorld.cpp edlib/src/edlib.cpp -o helloWorld -I edlib/include`.

If hello world was a C program, we would compile it like this:
```
    c++ -c edlib/src/edlib.cpp -o edlib.o -I edlib/include
    cc -c helloWorld.c -o helloWorld.o -I edlib/include
    c++ helloWorld.o edlib.o -o helloWorld
```

### Approach #2: Copying edlib header file and static library.
Instead of copying edlib source files, you could copy static library (check [Building](#building) on how to create static library). We also need to copy edlib header files. We get following project structure:
```
edlib/  -> copied from edlib
  include/
    edlib.h
  edlib.a
helloWorld.cpp -> your program
```

Now you can compile it with `c++ helloWorld.cpp -o helloWorld -I edlib/include -L edlib -ledlib`.

### Approach #3: Install edlib library on machine.
Alternatively, you could avoid copying any Edlib files and instead install libraries by running `sudo make install` (check [Building](#building) for exact instructions depending on approach you used for building). Now, all you have to do to compile your project is `c++ helloWorld.cpp -o helloWorld -ledlib`.
If you get error message like `cannot open shared object file: No such file or directory`, make sure that your linker includes path where edlib was installed.

## Development and contributing
Feel free to send pull requests and raise issues.

## Acknowledgements

### No one yet

## FAQ

### none yet