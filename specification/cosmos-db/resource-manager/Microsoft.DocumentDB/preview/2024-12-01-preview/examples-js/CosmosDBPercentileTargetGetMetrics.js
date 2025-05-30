const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieves the metrics determined by the given filter for the given account target region. This url is only for PBS and Replication Latency data
 *
 * @summary Retrieves the metrics determined by the given filter for the given account target region. This url is only for PBS and Replication Latency data
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBPercentileTargetGetMetrics.json
 */
async function cosmosDbDatabaseAccountRegionGetMetrics() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const targetRegion = "East US";
  const filter =
    "$filter=(name.value eq 'Probabilistic Bounded Staleness') and timeGrain eq duration'PT5M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.percentileTarget.listMetrics(
    resourceGroupName,
    accountName,
    targetRegion,
    filter,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
