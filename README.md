# Zi

![Go](https://github.com/Vitecoin/Zi/workflows/Go/badge.svg)

_A key value database._

Zi is a key value database originaly designed for storing coins and id's for Vitecoin. It can also be used for general storage with key value.

Zi should be put on a server or a local machine and in a directory with a file called `dump.zi`.

## Table of Contents

1. [Installing](#Install)
2. [Getting started](#Getting-started)
3. [Docker set-up](#Docker)
4. [Using Zi](#Using-Zi)

## Install

First things first your going to want to clone the repository.

```bash
git clone https://github.com/Vitecoin/Zi.git
```

Then all of install process is taken care of in the install file. You can just run:

```bash
chmod +x install
./install
```

This will install the CLI and get all of the dependencies.

## Getting started

The whole server and the database can be started with one simple command.

```bash
zi serve #by default the server will run on port 9090
```

You can specify a port with the following.

```bash
zi serve 5000
```

This will serve the server on port 5000.

You can add the `--background` flag to run the server as a daemon process.

```bash
zi serve 3000 --background
```

## Docker

You can spin up a Zi server in no time with Docker and Zi installed (Zi is optional but makes the process a bit easier).

```bash
sudo docker pull vitecoin/zi
```

This will install the zi docker cointainer.

Next if you don't have Zi installed you will want to run the docker container manually.

```bash
sudo docker run -p 5000:9090 -d --mount source=zi-presist,target=/app vitecoin/zi
```

This will run the docker container in the background on port 5000. You can also remove the `-d` flag to run as foreground process.

If you do have Zi installed run:

```bash
zi --docker --detached
```

This will run the container in the background.

```bash
zi --docker
```

This will run the docker container noramlly.

# Using Zi

Zi uses a purely REST based interface. Every action except setting up authentication can be done through the REST api. Setting up the REST api is really simple. First make sure the server is running either with docker or with the zi cli itself ([Read more](#Getting-started)). Once you have done this your ready start making requests!

## Writing Data

You can add data to the database with the following query.

```
http://[SERVER IP]:[PORT]/set?data=[JSON STRING WITH KEY AND VALUE]
```

The data url param takes in a json string with the following parameters:

```json
{
    "Key": "NAME OF ITEM",
    "Value": "ITEM VALUE "
}
```

_Capitalization does matter for the json keys._

## Reading Data

You can read data with the following:

```
http://[SERVER IP]:[PORT]/get?key=[KEY]
```

The get path takes in a parameter called key, which is the key the of an item which you want to get. If you are getting a binded key though meaning it is key that links to another database then you will want to pass in the following.

```
*[BINDED KEY NAME]:[KEY OF ITEM INSIDE BINDED DATABASE]
```

You can also read the entire database with the following:

```
http://[SERVER IP]:[PORT]/getall
```

This will return a json array of all the key value pairs and line numbers.

## Deleting Data

Data can be deleted with the following:

```
http://[SERVER IP]:[PORT]/del?key=[KEY]
```

This will delete all instances of that key.

## Binding Data

Two databases can binded or connected together using the bind action. Binding takes a key, which references another database it will use the data from the other database and store as the value of that key.
