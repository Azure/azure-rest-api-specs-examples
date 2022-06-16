const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getInformationAboutASnapshot() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const snapshotName = "mySnapshot";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.get(resourceGroupName, snapshotName);
  console.log(result);
}

getInformationAboutASnapshot().catch(console.error);
