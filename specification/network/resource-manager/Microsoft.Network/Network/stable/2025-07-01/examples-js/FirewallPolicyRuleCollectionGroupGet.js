const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary gets the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupGet.json
 */
async function getFirewallPolicyRuleCollectionGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyRuleCollectionGroups.get(
    "rg1",
    "firewallPolicy",
    "ruleCollectionGroup1",
  );
  console.log(result);
}
