const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the usages (most recent data) for the given database.
 *
 * @summary Retrieves the usages (most recent data) for the given database.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBDatabaseGetUsages.json
 */
async function cosmosDbDatabaseGetUsages() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseRid = "databaseRid";
  const filter = "$filter=name.value eq 'Storage'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.database.listUsages(
    resourceGroupName,
    accountName,
    databaseRid,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

cosmosDbDatabaseGetUsages().catch(console.error);
