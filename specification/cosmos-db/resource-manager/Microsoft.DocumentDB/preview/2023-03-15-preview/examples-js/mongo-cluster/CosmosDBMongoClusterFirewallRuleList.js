const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the firewall rules in a given mongo cluster.
 *
 * @summary List all the firewall rules in a given mongo cluster.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/mongo-cluster/CosmosDBMongoClusterFirewallRuleList.json
 */
async function listFirewallRulesOfTheMongoCluster() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "TestGroup";
  const mongoClusterName = "myMongoCluster";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.mongoClusters.listFirewallRules(
    resourceGroupName,
    mongoClusterName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
