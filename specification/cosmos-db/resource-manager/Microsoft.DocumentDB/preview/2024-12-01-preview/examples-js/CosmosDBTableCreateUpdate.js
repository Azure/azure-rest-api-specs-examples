const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update an Azure Cosmos DB Table
 *
 * @summary Create or update an Azure Cosmos DB Table
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBTableCreateUpdate.json
 */
async function cosmosDbTableReplace() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const tableName = "tableName";
  const createUpdateTableParameters = {
    location: "West US",
    options: {},
    resource: { id: "tableName" },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.tableResources.beginCreateUpdateTableAndWait(
    resourceGroupName,
    accountName,
    tableName,
    createUpdateTableParameters,
  );
  console.log(result);
}
