const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an Azure Cosmos DB SQL userDefinedFunction
 *
 * @summary Create or update an Azure Cosmos DB SQL userDefinedFunction
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-11-15/examples/CosmosDBSqlUserDefinedFunctionCreateUpdate.json
 */
async function cosmosDbSqlUserDefinedFunctionCreateUpdate() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const containerName = "containerName";
  const userDefinedFunctionName = "userDefinedFunctionName";
  const createUpdateSqlUserDefinedFunctionParameters = {
    options: {},
    resource: { body: "body", id: "userDefinedFunctionName" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginCreateUpdateSqlUserDefinedFunctionAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    containerName,
    userDefinedFunctionName,
    createUpdateSqlUserDefinedFunctionParameters,
  );
  console.log(result);
}
