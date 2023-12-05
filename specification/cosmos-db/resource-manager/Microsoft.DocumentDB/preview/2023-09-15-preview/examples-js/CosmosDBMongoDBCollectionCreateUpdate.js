const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an Azure Cosmos DB MongoDB Collection
 *
 * @summary Create or update an Azure Cosmos DB MongoDB Collection
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBMongoDBCollectionCreateUpdate.json
 */
async function cosmosDbMongoDbcollectionCreateUpdate() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rg1";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const collectionName = "collectionName";
  const createUpdateMongoDBCollectionParameters = {
    location: "West US",
    options: {},
    resource: {
      analyticalStorageTtl: 500,
      id: "collectionName",
      indexes: [
        {
          key: { keys: ["_ts"] },
          options: { expireAfterSeconds: 100, unique: true },
        },
        { key: { keys: ["_id"] } },
      ],
      shardKey: { testKey: "Hash" },
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginCreateUpdateMongoDBCollectionAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    collectionName,
    createUpdateMongoDBCollectionParameters
  );
  console.log(result);
}
