const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the properties of an existing Azure Cosmos DB Mongo User Definition with the given Id.
 *
 * @summary Retrieves the properties of an existing Azure Cosmos DB Mongo User Definition with the given Id.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBMongoDBUserDefinitionGet.json
 */
async function cosmosDbMongoDbuserDefinitionGet() {
  const subscriptionId = "mySubscriptionId";
  const mongoUserDefinitionId = "myMongoUserDefinitionId";
  const resourceGroupName = "myResourceGroupName";
  const accountName = "myAccountName";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.getMongoUserDefinition(
    mongoUserDefinitionId,
    resourceGroupName,
    accountName
  );
  console.log(result);
}

cosmosDbMongoDbuserDefinitionGet().catch(console.error);
