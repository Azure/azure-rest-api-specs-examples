const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of ElasticSan in a resource group.
 *
 * @summary gets a list of ElasticSan in a resource group.
 * x-ms-original-file: 2026-05-01-preview/ElasticSans_V2_ListByResourceGroup_MaximumSet_Gen.json
 */
async function elasticSansV2ListByResourceGroupMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.elasticSans.listByResourceGroup("resourcegroupname")) {
    resArray.push(item);
  }

  console.log(resArray);
}
