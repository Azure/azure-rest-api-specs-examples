const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function listSnapshots() {
  const subscriptionId = "subid1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.snapshots.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSnapshots().catch(console.error);
