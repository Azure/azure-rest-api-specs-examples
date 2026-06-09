const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists available object recommendations.
 *
 * @summary lists available object recommendations.
 * x-ms-original-file: 2026-04-01-preview/TuningOptionsListIndexRecommendations.json
 */
async function listAvailableIndexRecommendations() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.tuningOptionsOperations.listRecommendations(
    "exampleresourcegroup",
    "exampleserver",
    "index",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
