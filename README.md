HTTP Server Generator
&middot;
![GitHub release (latest by date)](https://img.shields.io/github/v/release/Peyton232/HTTPServerGenerator?display_name=tag)
<!-- [![Build status of the master branch on Linux/OSX](https://img.shields.io/travis/Martinsos/edlib/master?label=Linux%20%2F%20OSX%20build)](https://travis-ci.com/Martinsos/edlib)
[![Build status of the master branch on Windows](https://img.shields.io/appveyor/build/Martinsos/edlib/master?label=Windows%20build)](https://ci.appveyor.com/project/Martinsos/edlib/branch/master) -->
![GitHub issues](https://img.shields.io/github/issues-raw/Peyton232/HTTPServerGenerator)
![Issues Closed](https://img.shields.io/github/issues-closed/Peyton232/HTTPServerGenerator?display_name=tag)
![Open Milestones](https://img.shields.io/github/milestones/open/Peyton232/HTTPServerGenerator)
![Last Commit](https://img.shields.io/github/last-commit/Peyton232/HTTPServerGenerator)
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

### Approach #1: Directly pass in csv file thropugh env file


### Approach #2: Pass in csv file location thorugh cli argument 


### Approach #3: Turn on interpretation flag and pass in main.go of the server


## Development and contributing
Feel free to send pull requests and raise issues.

## Acknowledgements

### No one yet

## FAQ

### none yet
