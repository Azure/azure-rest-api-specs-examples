const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get an VolumeGroups.
 *
 * @summary Get an VolumeGroups.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/VolumeGroups_Get_MaximumSet_Gen.json
 */
async function volumeGroupsGetMaximumSetGen() {
  const subscriptionId = "aaaaaaaaaaaaaaaaaa";
  const resourceGroupName = "rgelasticsan";
  const elasticSanName = "ti7q-k952-1qB3J_5";
  const volumeGroupName = "u_5I_1j4t3";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const result = await client.volumeGroups.get(resourceGroupName, elasticSanName, volumeGroupName);
  console.log(result);
}

volumeGroupsGetMaximumSetGen().catch(console.error);
