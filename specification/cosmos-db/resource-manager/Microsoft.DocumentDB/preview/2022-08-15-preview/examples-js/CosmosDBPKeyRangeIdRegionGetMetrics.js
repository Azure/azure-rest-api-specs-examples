const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the metrics determined by the given filter for the given partition key range id and region.
 *
 * @summary Retrieves the metrics determined by the given filter for the given partition key range id and region.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBPKeyRangeIdRegionGetMetrics.json
 */
async function cosmosDbDatabaseAccountRegionGetMetrics() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const region = "West US";
  const databaseRid = "databaseRid";
  const collectionRid = "collectionRid";
  const partitionKeyRangeId = "0";
  const filter =
    "$filter=(name.value eq 'Max RUs Per Second') and timeGrain eq duration'PT1M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T23:58:55.2780000Z";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.partitionKeyRangeIdRegion.listMetrics(
    resourceGroupName,
    accountName,
    region,
    databaseRid,
    collectionRid,
    partitionKeyRangeId,
    filter
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

cosmosDbDatabaseAccountRegionGetMetrics().catch(console.error);
