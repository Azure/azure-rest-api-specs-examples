const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete an VolumeGroup.
 *
 * @summary delete an VolumeGroup.
 * x-ms-original-file: 2025-09-01/VolumeGroups_Delete_MinimumSet_Gen.json
 */
async function volumeGroupsDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  await client.volumeGroups.delete("resourcegroupname", "elasticsanname", "volumegroupname");
}
