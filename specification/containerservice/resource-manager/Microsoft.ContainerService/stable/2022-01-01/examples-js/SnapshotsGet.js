const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function getSnapshot() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "snapshot1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.snapshots.get(resourceGroupName, resourceName);
  console.log(result);
}

getSnapshot().catch(console.error);
