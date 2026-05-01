const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a snapshot.
 *
 * @summary creates or updates a snapshot.
 * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_Create_FromAnElasticSanVolumeSnapshot.json
 */
async function createASnapshotFromAnElasticSanVolumeSnapshot() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.snapshots.createOrUpdate("myResourceGroup", "mySnapshot", {
    location: "West US",
    creationData: {
      createOption: "CopyFromSanSnapshot",
      elasticSanResourceId:
        "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ElasticSan/elasticSans/myElasticSan/volumegroups/myElasticSanVolumeGroup/snapshots/myElasticSanVolumeSnapshot",
    },
  });
  console.log(result);
}
