const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieves metric definitions for the given collection.
 *
 * @summary Retrieves metric definitions for the given collection.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBCollectionGetMetricDefinitions.json
 */
async function cosmosDbCollectionGetMetricDefinitions() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseRid = "databaseRid";
  const collectionRid = "collectionRid";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.collection.listMetricDefinitions(
    resourceGroupName,
    accountName,
    databaseRid,
    collectionRid,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
