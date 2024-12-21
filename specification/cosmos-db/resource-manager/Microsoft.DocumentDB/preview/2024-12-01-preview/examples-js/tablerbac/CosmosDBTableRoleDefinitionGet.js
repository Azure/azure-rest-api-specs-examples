const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieves the properties of an existing Azure Cosmos DB Table Role Definition with the given Id.
 *
 * @summary Retrieves the properties of an existing Azure Cosmos DB Table Role Definition with the given Id.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/tablerbac/CosmosDBTableRoleDefinitionGet.json
 */
async function cosmosDbTableRoleDefinitionGet() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "myResourceGroupName";
  const accountName = "myAccountName";
  const roleDefinitionId = "myRoleDefinitionId";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.tableResources.getTableRoleDefinition(
    resourceGroupName,
    accountName,
    roleDefinitionId,
  );
  console.log(result);
}
