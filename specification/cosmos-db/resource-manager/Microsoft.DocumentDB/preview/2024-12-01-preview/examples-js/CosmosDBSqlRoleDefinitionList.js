const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieves the list of all Azure Cosmos DB SQL Role Definitions.
 *
 * @summary Retrieves the list of all Azure Cosmos DB SQL Role Definitions.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBSqlRoleDefinitionList.json
 */
async function cosmosDbSqlRoleDefinitionList() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "mySubscriptionId";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "myResourceGroupName";
  const accountName = "myAccountName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sqlResources.listSqlRoleDefinitions(
    resourceGroupName,
    accountName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
