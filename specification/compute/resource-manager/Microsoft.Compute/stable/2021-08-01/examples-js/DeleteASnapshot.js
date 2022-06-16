const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteASnapshot() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const snapshotName = "mySnapshot";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.beginDeleteAndWait(resourceGroupName, snapshotName);
  console.log(result);
}

deleteASnapshot().catch(console.error);
