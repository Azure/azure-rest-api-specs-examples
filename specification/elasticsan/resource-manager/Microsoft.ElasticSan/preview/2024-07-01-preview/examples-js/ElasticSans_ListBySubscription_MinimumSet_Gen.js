const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets a list of ElasticSans in a subscription
 *
 * @summary Gets a list of ElasticSans in a subscription
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/ElasticSans_ListBySubscription_MinimumSet_Gen.json
 */
async function elasticSansListBySubscriptionMinimumSetGen() {
  const subscriptionId = process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "subscriptionid";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.elasticSans.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
