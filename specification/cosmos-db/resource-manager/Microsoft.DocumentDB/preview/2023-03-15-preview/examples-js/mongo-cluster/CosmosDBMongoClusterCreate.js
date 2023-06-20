const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 *
 * @summary Create or update a mongo cluster. Update overwrites all properties for the resource. To only modify some of the properties, use PATCH.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/mongo-cluster/CosmosDBMongoClusterCreate.json
 */
async function createANewMongoCluster() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "TestResourceGroup";
  const mongoClusterName = "myMongoCluster";
  const parameters = {
    administratorLogin: "mongoAdmin",
    administratorLoginPassword: "password",
    location: "westus2",
    nodeGroupSpecs: [
      {
        diskSizeGB: 128,
        enableHa: true,
        kind: "Shard",
        nodeCount: 3,
        sku: "M30",
      },
    ],
    serverVersion: "5.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    mongoClusterName,
    parameters
  );
  console.log(result);
}
