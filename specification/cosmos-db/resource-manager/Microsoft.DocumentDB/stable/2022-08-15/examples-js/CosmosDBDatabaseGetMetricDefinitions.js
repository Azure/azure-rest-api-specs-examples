const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves metric definitions for the given database.
 *
 * @summary Retrieves metric definitions for the given database.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBDatabaseGetMetricDefinitions.json
 */
async function cosmosDbDatabaseGetMetricDefinitions() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseRid = "databaseRid";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.database.listMetricDefinitions(
    resourceGroupName,
    accountName,
    databaseRid
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

cosmosDbDatabaseGetMetricDefinitions().catch(console.error);
