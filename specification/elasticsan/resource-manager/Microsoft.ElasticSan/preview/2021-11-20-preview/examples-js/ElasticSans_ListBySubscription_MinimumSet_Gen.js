const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of ElasticSans in a subscription
 *
 * @summary Gets a list of ElasticSans in a subscription
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/ElasticSans_ListBySubscription_MinimumSet_Gen.json
 */
async function elasticSansListBySubscriptionMinimumSetGen() {
  const subscriptionId = "aaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.elasticSans.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

elasticSansListBySubscriptionMinimumSetGen().catch(console.error);
