const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a mongo cluster firewall rule.
 *
 * @summary Deletes a mongo cluster firewall rule.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/mongo-cluster/CosmosDBMongoClusterFirewallRuleDelete.json
 */
async function deleteTheFirewallRuleOfTheMongoCluster() {
  const subscriptionId =
    process.env["COSMOSDB_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "TestGroup";
  const mongoClusterName = "myMongoCluster";
  const firewallRuleName = "rule1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.mongoClusters.beginDeleteFirewallRuleAndWait(
    resourceGroupName,
    mongoClusterName,
    firewallRuleName
  );
  console.log(result);
}
