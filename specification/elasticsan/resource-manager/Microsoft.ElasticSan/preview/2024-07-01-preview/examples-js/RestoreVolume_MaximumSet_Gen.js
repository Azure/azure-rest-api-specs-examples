const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Restore Soft Deleted Volumes. The volume name is obtained by using the API to list soft deleted volumes by volume group
 *
 * @summary Restore Soft Deleted Volumes. The volume name is obtained by using the API to list soft deleted volumes by volume group
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/RestoreVolume_MaximumSet_Gen.json
 */
async function restoreVolumeMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const volumeGroupName = "volumegroupname";
  const volumeName = "volumename-1741526907";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.beginRestoreVolumeAndWait(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    volumeName,
  );
  console.log(result);
}
