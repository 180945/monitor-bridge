const ABI = require("./abi.json");
require("dotenv").config();
const ethers = require('ethers');
const {CronJob} = require('cron');

const provider = new ethers.JsonRpcProvider("https://rpc.supersonic.bvm.network");
const signer = new ethers.Wallet(process.env.PRIVATE_KEY || "", provider);
const contract = new ethers.Contract("0xdcae67621fB30C0a9FC75576e5c7d42F988d2B2B", ABI, provider);

async function sendTx () {
    const tx = await signer.sendTransaction({
        to: '0x9699b31b25D71BDA4819bBe66244E9130cEE62b7',
        value: ethers.parseUnits('0.0000001', 'ether'),
    });
    console.log(tx);
}

async function filterEvent(startBlock) {
    const latestBlock = await provider.getBlockNumber();
    console.log({startBlock});
    console.log({latestBlock});

    if (latestBlock === startBlock) {
        return startBlock;
    }

    if (latestBlock - startBlock >= 1000) {
        await sendTx();
        startBlock = latestBlock;
        return startBlock;
    }

    // scan event
    const execute =  await contract.queryFilter("BlockExecution", startBlock, latestBlock);
    if (execute != null && execute.length > 0) {
        await sendTx();
    }

    const commit =  await contract.queryFilter("BlockCommit", startBlock, latestBlock);
    if (commit != null && commit.length > 0) {
        await sendTx();
    }

    return latestBlock;
}

(function main() {
    let startBlock = 461;
    const job = new CronJob(
        '*/20 * * * * *', // cronTime
        async function () {
            startBlock = await filterEvent(startBlock);
        }, // onTick
        null, // onComplete
        false, // start
        'America/Los_Angeles' // timeZone
    );
    job.start();
})();

// block execution
// contract.on("BlockExecution", (blockNumber, blockHash, commitment, event)=>{
//     let executeEvent ={
//         blockNumber: blockNumber,
//         blockHash: blockHash,
//         commitment: commitment,
//         eventData: event,
//     }
//     console.log(JSON.stringify(executeEvent, null, 4))
// })
//
// // block commit
// contract.on("BlockCommit", (batchNumber, batchHash, commitment, event)=>{
//     let commitEvent ={
//         batchNumber: batchNumber,
//         batchHash: batchHash,
//         commitment: commitment,
//         eventData: event,
//     }
//     console.log(JSON.stringify(commitEvent, null, 4))
// })