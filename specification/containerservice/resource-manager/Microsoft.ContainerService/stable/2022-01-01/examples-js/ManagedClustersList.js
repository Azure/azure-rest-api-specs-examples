const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function listManagedClusters() {
  const subscriptionId = "subid1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedClusters.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedClusters().catch(console.error);
