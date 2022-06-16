const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createASnapshotFromAnExistingSnapshotInTheSameOrADifferentSubscriptionInADifferentRegion() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const snapshotName = "mySnapshot2";
  const snapshot = {
    creationData: {
      createOption: "CopyStart",
      sourceResourceId:
        "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1",
    },
    location: "West US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.beginCreateOrUpdateAndWait(
    resourceGroupName,
    snapshotName,
    snapshot
  );
  console.log(result);
}

createASnapshotFromAnExistingSnapshotInTheSameOrADifferentSubscriptionInADifferentRegion().catch(
  console.error
);
