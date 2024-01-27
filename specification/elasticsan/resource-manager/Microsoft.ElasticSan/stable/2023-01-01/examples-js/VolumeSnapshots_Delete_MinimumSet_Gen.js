const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Volume Snapshot.
 *
 * @summary Delete a Volume Snapshot.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Delete_MinimumSet_Gen.json
 */
async function volumeSnapshotsDeleteMinimumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const volumeGroupName = "volumegroupname";
  const snapshotName = "snapshotname";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeSnapshots.beginDeleteAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    snapshotName,
  );
  console.log(result);
}
