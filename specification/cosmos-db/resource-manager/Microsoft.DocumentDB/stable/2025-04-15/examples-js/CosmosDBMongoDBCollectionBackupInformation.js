const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieves continuous backup information for a Mongodb collection.
 *
 * @summary Retrieves continuous backup information for a Mongodb collection.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBMongoDBCollectionBackupInformation.json
 */
async function cosmosDbMongoDbcollectionBackupInformation() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rgName";
  const accountName = "ddb1";
  const databaseName = "databaseName";
  const collectionName = "collectionName";
  const location = {
    location: "North Europe",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoDBResources.beginRetrieveContinuousBackupInformationAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    collectionName,
    location,
  );
  console.log(result);
}
