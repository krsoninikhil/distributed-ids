# Distributed IDs

## Requirements

- A GNU/Linux machine (tested on Ubuntu 16.04)
- Install cURL command line utility
```bash
sudo apt-get install curl
```
- Install Docker using:
```bash
sudo apt-get install docker
```
- Install Node.js using:
```bash
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.8/install.sh | bash
```
The script clones the nvm repository to ~/.nvm and adds the source line to your profile (~/.bash_profile, ~/.zshrc, ~/.profile, or ~/.bashrc).
```bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm
```
```bash
nvm install 6.10.0
nvm use 6.10.0
```
Installing Node.js will also install NPM, however it is recommended that you confirm the version of NPM installed.

## Run Application

- Download or clone this repository and change directory to `fabcar` example:
```bash
cd distributed-ids/fabcar
```
- Run the following command to install the Fabric dependencies for the applications:
```bash
npm install
```
- Start Hyperledger Fabric using following script
```bash
./startFabric.sh
```
- Once the application is started, change directory to `node_project` to interact with ledger using it's Node SDK
```bash
cd ../node_project
```
- Start user interface web application server:
```bash
node server.js
```
- To register user, go to `http://localhost/login.html`
- To save all information of user `http://localhost/process_get`

## License

Hyperledger Fabric [license](https://github.com/hyperledger/fabric-samples/blob/release/LICENSE)
