// https://wsrpc.supersonic.bvm.network/
const ABI = require("./abi.json");
require("dotenv").config();
const ethers = require('ethers');

const provider = new ethers.JsonRpcProvider("https://wsrpc.supersonic.bvm.network/");
const signer = new ethers.Wallet(process.env.PRIVATE_KEY || "", provider);

async function sendTx () {
    const tx = await signer.sendTransaction({
        to: '0x9699b31b25D71BDA4819bBe66244E9130cEE62b7',
        value: ethers.parseUnits('0.000001', 'ether'),
    });
    console.log(tx);
}

(async function main(){
    console.log("dmc");

    const contract = new ethers.Contract("0xdcae67621fB30C0a9FC75576e5c7d42F988d2B2B", ABI, provider);

    // block execution
    contract.on("BlockExecution", (blockNumber, blockHash, commitment, event)=>{
        let executeEvent ={
            blockNumber: blockNumber,
            blockHash: blockHash,
            commitment: commitment,
            eventData: event,
        }
        console.log(JSON.stringify(executeEvent, null, 4));
        sendTx();
    })

    // block commit
    contract.on("BlockCommit", (batchNumber, batchHash, commitment, event)=>{
        let commitEvent ={
            batchNumber: batchNumber,
            batchHash: batchHash,
            commitment: commitment,
            eventData: event,
        }
        console.log(JSON.stringify(commitEvent, null, 4));
        sendTx();
    })
})();