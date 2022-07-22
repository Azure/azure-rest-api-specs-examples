const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Azure Cosmos DB SQL userDefinedFunction.
 *
 * @summary Deletes an existing Azure Cosmos DB SQL userDefinedFunction.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBSqlUserDefinedFunctionDelete.json
 */
async function cosmosDbSqlUserDefinedFunctionDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const containerName = "containerName";
  const userDefinedFunctionName = "userDefinedFunctionName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginDeleteSqlUserDefinedFunctionAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    containerName,
    userDefinedFunctionName
  );
  console.log(result);
}

cosmosDbSqlUserDefinedFunctionDelete().catch(console.error);
