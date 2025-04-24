const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List Volumes in a VolumeGroup.
 *
 * @summary List Volumes in a VolumeGroup.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/Volumes_ListByVolumeGroup_MaximumSet_Gen.json
 */
async function volumesListByVolumeGroupMaximumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const resourceGroupName = process.env["ELASTICSANS_RESOURCE_GROUP"] || "resourcegroupname";
  const elasticSanName = "elasticsanname";
  const volumeGroupName = "volumegroupname";
  const xMsAccessSoftDeletedResources = "true";
  const options = {
    xMsAccessSoftDeletedResources,
  };
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.volumes.listByVolumeGroup(
    resourceGroupName,
    elasticSanName,
    volumeGroupName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
