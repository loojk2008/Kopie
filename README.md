# Kopie

[![Build Status](https://travis-ci.com/CIP-NL/Kopie.svg?branch=master)](https://travis-ci.com/CIP-NL/Kopie)
[![Go Report Card](https://goreportcard.com/badge/github.com/CIP-NL/Kopie)](https://goreportcard.com/report/github.com/CIP-NL/Kopie)

Go based data replication between postgres DB's when logical or streaming replication is not suitable.

Simply compile kopie.go, create a configuration toml file, and start the binary with and environment variable
KOPIE_CONFIG_PATH=path/to/config.toml

Currently the only supported protocol is pump:

```toml
[[protocols]]
# a unique identifier    
  name = "mainProcedure"
# procedure name  
  type = "pump"  
  [protocols.pump]
# table we start replicated  
   master = "kopie_test"
   tables = [
        "test",
        "test2",
   ]
# replication destination   
   slave = "kopie_test2"
# create new tables if schema has changed   
   automigrate = true
# sleeping time in seconds between replications   
   period = 1

```