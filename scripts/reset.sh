#!/bin/bash

BINARY1=cosmoverse-workshopd
BINARY2=simple-dexd
CHAIN_DIR1=./cosmoverse-workshop
CHAIN_DIR2=./simple-dex
CHAINID_1=test-1
CHAINID_2=test-2

echo "Removing previous data..."
rm -rf $CHAIN_DIR1 &> /dev/null
rm -rf $CHAIN_DIR2 &> /dev/null

