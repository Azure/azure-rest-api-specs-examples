const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a mongo cluster.
 *
 * @summary Deletes a mongo cluster.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/mongo-cluster/CosmosDBMongoClusterDelete.json
 */
async function deleteTheMongoCluster() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "TestResourceGroup";
  const mongoClusterName = "myMongoCluster";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.beginDeleteAndWait(resourceGroupName, mongoClusterName);
  console.log(result);
}
