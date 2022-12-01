const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates tags on a managed cluster snapshot.
 *
 * @summary Updates tags on a managed cluster snapshot.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-10-02-preview/examples/ManagedClusterSnapshotsUpdateTags.json
 */
async function updateManagedClusterSnapshotTags() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "snapshot1";
  const parameters = { tags: { key2: "new-val2", key3: "val3" } };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusterSnapshots.updateTags(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

updateManagedClusterSnapshotTags().catch(console.error);
