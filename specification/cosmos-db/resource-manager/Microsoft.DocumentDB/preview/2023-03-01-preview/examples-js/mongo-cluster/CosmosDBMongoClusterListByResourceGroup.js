const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the mongo clusters in a given resource group.
 *
 * @summary List all the mongo clusters in a given resource group.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/mongo-cluster/CosmosDBMongoClusterListByResourceGroup.json
 */
async function listTheMongoClustersByResourceGroup() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "TestResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.mongoClusters.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
