# Cosmoverse 2022 Workshop 

**eightball** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Usecase

User will submit a token offering, and receive back a fortune.

Chain 1 is a eightball which gives out fortunes if you give it an offering in any denom but actually only wants a specific denom

Chain 2 is a simple dex, which will take the token offering and swap it into the denom that the eightball wants


![alt text](/img/demo_workflow.png)

## CLI commands

```
rm -rf ~/.cosmoverse-workshop && rm -rf ~/.simple-dex && hermes keys delete --chain simpledex --all && hermes keys delete --chain cosmoverseworkshop --all

ignite chain serve

hermes keys add --chain cosmoverseworkshop --mnemonic-file ~/.hermes/cvworkshop-mnemonic.txt && hermes keys add --chain simpledex --mnemonic-file ~/.hermes/simpledex-mnemonic.txt

hermes create connection --a-chain cosmoverseworkshop --b-chain simpledex

cosmoverse-workshopd q ibc connection connections 
simple-dexd q ibc connection connections        

hermes start --full-scan

cosmoverse-workshopd tx eightball connect-to-dex connection-0 --from alice --chain-id cosmoverseworkshop --gas auto 

cosmoverse-workshopd q ibc channel channels
simple-dexd q ibc channel channels

cosmoverse-workshopd tx eightball feeling-lucky 50stake --from alice --chain-id cosmoverseworkshop --gas auto 

cosmoverse-workshopd tx eightball feeling-lucky 100stake --from bob --chain-id cosmoverseworkshop --gas auto

cosmoverse-workshopd q bank balances

cosmoverse-workshopd q eightball list-fortunes
cosmoverse-workshopd q eightball show-fortune <address>

```

