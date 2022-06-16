const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateASnapshotWithAcceleratedNetworking() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const snapshotName = "mySnapshot";
  const snapshot = {
    diskSizeGB: 20,
    supportedCapabilities: { acceleratedNetwork: false },
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

updateASnapshotWithAcceleratedNetworking().catch(console.error);
