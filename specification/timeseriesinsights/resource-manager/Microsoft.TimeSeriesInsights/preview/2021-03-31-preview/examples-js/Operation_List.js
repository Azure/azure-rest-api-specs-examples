const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Lists all of the available Time Series Insights related operations.
 *
 * @summary Lists all of the available Time Series Insights related operations.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/Operation_List.json
 */
async function listAvailableOperationsForTheTimeSeriesInsightsResourceProvider() {
  const subscriptionId =
    process.env["TIMESERIESINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

async function main() {
  listAvailableOperationsForTheTimeSeriesInsightsResourceProvider();
}

main().catch(console.error);
