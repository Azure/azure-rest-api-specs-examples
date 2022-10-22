const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List VolumeGroups.
 *
 * @summary List VolumeGroups.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/VolumeGroups_ListByElasticSan_MaximumSet_Gen.json
 */
async function volumeGroupsListByElasticSanMaximumSetGen() {
  const subscriptionId = "aaaaaaaaaaaaaaaaaa";
  const resourceGroupName = "rgelasticsan";
  const elasticSanName = "ti7q-k952-1qB3J_5";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.volumeGroups.listByElasticSan(resourceGroupName, elasticSanName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

volumeGroupsListByElasticSanMaximumSetGen().catch(console.error);
