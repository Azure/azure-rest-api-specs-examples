const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of ElasticSan in a resource group.
 *
 * @summary gets a list of ElasticSan in a resource group.
 * x-ms-original-file: 2025-09-01/ElasticSans_ListByResourceGroup_MaximumSet_Gen.json
 */
async function elasticSansListByResourceGroupMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionid";
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.elasticSans.listByResourceGroup("resourcegroupname")) {
    resArray.push(item);
  }

  console.log(resArray);
}
