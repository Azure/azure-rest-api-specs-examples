const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the SQL userDefinedFunction under an existing Azure Cosmos DB database account.
 *
 * @summary Lists the SQL userDefinedFunction under an existing Azure Cosmos DB database account.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBSqlUserDefinedFunctionList.json
 */
async function cosmosDbSqlUserDefinedFunctionList() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rgName";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const containerName = "containerName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlResources.listSqlUserDefinedFunctions(
    resourceGroupName,
    accountName,
    databaseName,
    containerName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
