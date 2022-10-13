const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Azure Cosmos DB Mongo User Definition.
 *
 * @summary Creates or updates an Azure Cosmos DB Mongo User Definition.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBMongoDBUserDefinitionCreateUpdate.json
 */
async function cosmosDbMongoDbuserDefinitionCreateUpdate() {
  const subscriptionId = "mySubscriptionId";
  const mongoUserDefinitionId = "myMongoUserDefinitionId";
  const resourceGroupName = "myResourceGroupName";
  const accountName = "myAccountName";
  const createUpdateMongoUserDefinitionParameters = {
    customData: "My custom data",
    databaseName: "sales",
    mechanisms: "SCRAM-SHA-256",
    password: "myPassword",
    roles: [{ db: "sales", role: "myReadRole" }],
    userName: "myUserName",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginCreateUpdateMongoUserDefinitionAndWait(
    mongoUserDefinitionId,
    resourceGroupName,
    accountName,
    createUpdateMongoUserDefinitionParameters
  );
  console.log(result);
}

cosmosDbMongoDbuserDefinitionCreateUpdate().catch(console.error);
