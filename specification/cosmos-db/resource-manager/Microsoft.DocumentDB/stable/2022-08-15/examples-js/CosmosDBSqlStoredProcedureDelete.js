const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Azure Cosmos DB SQL storedProcedure.
 *
 * @summary Deletes an existing Azure Cosmos DB SQL storedProcedure.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBSqlStoredProcedureDelete.json
 */
async function cosmosDbSqlStoredProcedureDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const containerName = "containerName";
  const storedProcedureName = "storedProcedureName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginDeleteSqlStoredProcedureAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    containerName,
    storedProcedureName
  );
  console.log(result);
}

cosmosDbSqlStoredProcedureDelete().catch(console.error);
