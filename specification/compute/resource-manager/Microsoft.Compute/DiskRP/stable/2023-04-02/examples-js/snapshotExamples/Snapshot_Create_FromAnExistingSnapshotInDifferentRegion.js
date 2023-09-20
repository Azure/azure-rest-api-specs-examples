const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a snapshot.
 *
 * @summary Creates or updates a snapshot.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/snapshotExamples/Snapshot_Create_FromAnExistingSnapshotInDifferentRegion.json
 */
async function createASnapshotFromAnExistingSnapshotInTheSameOrADifferentSubscriptionInADifferentRegion() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
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
