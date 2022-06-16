const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateASnapshot() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const snapshotName = "mySnapshot";
  const snapshot = {
    diskSizeGB: 20,
    tags: { department: "Development", project: "UpdateSnapshots" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.beginUpdateAndWait(
    resourceGroupName,
    snapshotName,
    snapshot
  );
  console.log(result);
}

updateASnapshot().catch(console.error);
