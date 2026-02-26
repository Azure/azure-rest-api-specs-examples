const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Volume Group.
 *
 * @summary create a Volume Group.
 * x-ms-original-file: 2025-09-01/VolumeGroups_Create_MinimumSet_Gen.json
 */
async function volumeGroupsCreateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeGroups.create(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
    {},
  );
  console.log(result);
}
