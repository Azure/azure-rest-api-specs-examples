const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a Volume Snapshot.
 *
 * @summary delete a Volume Snapshot.
 * x-ms-original-file: 2025-09-01/VolumeSnapshots_Delete_MinimumSet_Gen.json
 */
async function volumeSnapshotsDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  await client.volumeSnapshots.delete(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
    "snapshotname",
  );
}
