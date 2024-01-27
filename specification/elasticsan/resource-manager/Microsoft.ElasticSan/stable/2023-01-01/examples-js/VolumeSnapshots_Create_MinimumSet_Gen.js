const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Volume Snapshot.
 *
 * @summary Create a Volume Snapshot.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Create_MinimumSet_Gen.json
 */
async function volumeSnapshotsCreateMinimumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const volumeGroupName = "volumegroupname";
  const snapshotName = "snapshotname";
  const parameters = {
    properties: {
      creationData: {
        sourceId:
          "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeSnapshots.beginCreateAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    snapshotName,
    parameters,
  );
  console.log(result);
}
