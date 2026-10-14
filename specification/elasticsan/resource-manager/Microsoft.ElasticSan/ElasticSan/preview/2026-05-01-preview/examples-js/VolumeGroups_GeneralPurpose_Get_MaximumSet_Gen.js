const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get an VolumeGroups.
 *
 * @summary get an VolumeGroups.
 * x-ms-original-file: 2026-05-01-preview/VolumeGroups_GeneralPurpose_Get_MaximumSet_Gen.json
 */
async function volumeGroupsGeneralPurposeGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeGroups.get(
    "resourcegroupname",
    "elasticsanname",
    "volumegroupname",
  );
  console.log(result);
}
