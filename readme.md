# Cosmoverse 2022 Workshop

**8ball** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).



rm -rf ~/.cosmoverse-workshop && rm -rf ~/.simple-dex && hermes keys delete --chain simpledex --all && hermes keys delete --chain cosmoverseworkshop --all

hermes keys add --chain cosmoverseworkshop --mnemonic-file ~/.hermes/cvworkshop-mnemonic.txt && hermes keys add --chain simpledex --mnemonic-file ~/.hermes/simpledex-mnemonic.txt

hermes create connection --a-chain cosmoverseworkshop --b-chain simpledex

cosmoverse-workshopd q ibc connection connections 
simple-dexd q ibc connection connections        

hermes start --full-scan

cosmoverse-workshopd tx eightball connect-to-dex connection-0 --from alice --chain-id cosmoverseworkshop --gas auto 
cosmoverse-workshopd q ibc channel channels
cosmoverse-workshopd tx eightball feeling-lucky 50stake --from alice --chain-id cosmoverseworkshop  
cosmoverse-workshopd tx eightball feeling-lucky 100stake --from bob --chain-id cosmoverseworkshop 

cosmoverse-workshopd q bank balances

cosmoverse-workshopd q eightball list-fortunes
cosmoverse-workshopd q eightball show-fortune


