# Robot Simulator
This is an application to simulate toy robot moving on a square tabletop, The robot is free to roam around the surface of the table, but must be prevented from falling to destruction. Any movement that would result in the robot falling from the table must be prevented, however further valid movement commands must still be allowed.

Table of Contents
=================

  * [Preparations](#preparations)
    * [Install GoLang](#install-golang)
    * [Set Up ENV](#set-up-env)
    * [Go Version](#go-version)
    * [Working Directory](#working-directory)
  * [Running tests](#running-tests)
  * [Compiling/Running the Application](#running-the-application)
  * [Commands](#commands)

## Preparations

### Install Golang

Please Follow the instruction [here for mac](https://golang.org/doc/install)


### Set Up ENV

Make sure `$GOPATH` environment variable are correctly setup

    #this should be your workspace for all the go packages (result will be different from yours)
    $ echo $GOPATH
    $ /Users/macuser/Documents/gowork

### GO version

At the time of writing, we are using `1.7.4` .
to check your version

    $ go version
    $ go version go1.7.4 darwin/amd64

### Working Directory

Ensure the structure of your working directory is `$GOPATH/src/robot-simulator/.`, copy or clone all the source code inside this working directory.

## Running tests

## Compiling/Running the Application

### Compiling the Application

To Compile the code first make sure you are in the working director

  $ pwd
  $ /Users/macuser/Document/gowork/src/robot-simulator

Then download all dependencies

  $ go get -t -v ./...

Then compile the code, which will create an executable under `$GOPATH/bin/robot-simulator`

  $ go build *.go

### Running the Application

You can run the application from the executable file generated after you compile the code

  $ /Users/macuser/Document/gowork/bin/robot-simulator
  or
  $ /Users/macuser/Document/gowork/src/robot-simulator/main

Alternatively you can also run the code without compiling the code from your working directory

  $ go run *.go

## Commands

Below are available commands to simulate the robot:
  `PLACE X,Y,F`
  PLACE will put the toy robot on the table in position X,Y and facing NORTH,
  SOUTH, EAST or WEST.

  `MOVE`
  MOVE will move the toy robot one unit forward in the direction it is
  currently facing.

  `LEFT` & `RIGHT`
  LEFT and RIGHT will rotate the robot 90 degrees in the specified direction
  without changing the position of the robot.

  `REPORT`
  REPORT will announce the X,Y and F of the robot. This can be in any form,
  but standard output is sufficient.
