const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets information about a DocumentDB cluster firewall rule.
 *
 * @summary gets information about a DocumentDB cluster firewall rule.
 * x-ms-original-file: 2026-09-01/MongoClusters_FirewallRuleGet.json
 */
async function getsAFirewallRuleOnADocumentDBClusterResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.get("TestGroup", "myMongoCluster", "rule1");
  console.log(result);
}
