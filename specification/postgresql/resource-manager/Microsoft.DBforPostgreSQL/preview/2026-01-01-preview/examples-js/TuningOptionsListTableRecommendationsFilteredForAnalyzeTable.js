const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists available object recommendations.
 *
 * @summary lists available object recommendations.
 * x-ms-original-file: 2026-01-01-preview/TuningOptionsListTableRecommendationsFilteredForAnalyzeTable.json
 */
async function listAvailableTableRecommendationsFilteredToExclusivelyGetThoseOfAnalyzeTableType() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.tuningOptions.listRecommendations(
    "exampleresourcegroup",
    "exampleserver",
    "table",
    { recommendationType: "AnalyzeTable" },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
