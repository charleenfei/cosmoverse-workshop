# Cosmoverse 2022 Workshop 
**8ball** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/charleenfei/cosmoverse-workshop@latest! | sudo bash
```
`charleenfei/cosmoverse-workshop` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

rm -rf ~/.cosmoverse-workshop && rm -rf ~/.simple-dex && hermes keys delete --chain simpledex --all && hermes keys delete --chain cosmoverseworkshop --all

hermes keys add --chain cosmoverseworkshop --mnemonic-file ~/.hermes/cvworkshop-mnemonic.txt && hermes keys add --chain simpledex --mnemonic-file ~/.hermes/simpledex-mnemonic.txt

hermes create connection --a-chain cosmoverseworkshop --b-chain simpledex

cosmoverse-workshopd q ibc connection connections 
simple-dexd q ibc connection connections        

hermes start --full-scan

cosmoverse-workshopd tx eightball connect-to-dex connection-0 --from alice --chain-id cosmoverseworkshop --gas auto 
cosmoverse-workshopd q ibc channel channels
cosmoverse-workshopd tx eightball feeling-lucky 50stake --from alice --chain-id cosmoverseworkshop --gas auto 
cosmoverse-workshopd tx eightball feeling-lucky 100stake --from bob --chain-id cosmoverseworkshop --gas auto

cosmoverse-workshopd q bank balances

cosmoverse-workshopd q eightball list-fortunes
cosmoverse-workshopd q eightball show-fortune


