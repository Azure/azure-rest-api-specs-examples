const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing mongo cluster. The request body can contain one to many of the properties present in the normal mongo cluster definition.
 *
 * @summary Updates an existing mongo cluster. The request body can contain one to many of the properties present in the normal mongo cluster definition.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/mongo-cluster/CosmosDBMongoClusterUpdate.json
 */
async function updateTheMongoCluster() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "TestResourceGroup";
  const mongoClusterName = "myMongoCluster";
  const parameters = {
    administratorLogin: "mongoAdmin",
    administratorLoginPassword: "password",
    nodeGroupSpecs: [
      {
        diskSizeGB: 256,
        enableHa: true,
        kind: "Shard",
        nodeCount: 4,
        sku: "M50",
      },
    ],
    serverVersion: "5.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.beginUpdateAndWait(
    resourceGroupName,
    mongoClusterName,
    parameters,
  );
  console.log(result);
}
