const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update an VolumeGroup.
 *
 * @summary update an VolumeGroup.
 * x-ms-original-file: 2025-09-01/VolumeGroups_Update_MinimumSet_Gen.json
 */
async function volumeGroupsUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeGroups.update(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
    {},
  );
  console.log(result);
}
