const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing Azure Cosmos DB Mongo User Definition.
 *
 * @summary Deletes an existing Azure Cosmos DB Mongo User Definition.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBMongoDBUserDefinitionDelete.json
 */
async function cosmosDbMongoDbuserDefinitionDelete() {
  const subscriptionId = "mySubscriptionId";
  const mongoUserDefinitionId = "myMongoUserDefinitionId";
  const resourceGroupName = "myResourceGroupName";
  const accountName = "myAccountName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginDeleteMongoUserDefinitionAndWait(
    mongoUserDefinitionId,
    resourceGroupName,
    accountName
  );
  console.log(result);
}

cosmosDbMongoDbuserDefinitionDelete().catch(console.error);
