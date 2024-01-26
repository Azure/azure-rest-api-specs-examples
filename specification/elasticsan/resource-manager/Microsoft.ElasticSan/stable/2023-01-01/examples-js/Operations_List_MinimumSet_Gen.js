const { ElasticSanManagement } = require("@azure/arm-elasticsan");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of ElasticSan operations.
 *
 * @summary Gets a list of ElasticSan operations.
 * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/Operations_List_MinimumSet_Gen.json
 */
async function operationsListMinimumSetGen() {
  const subscriptionId =
    process.env["ELASTICSANS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ElasticSanManagement(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
