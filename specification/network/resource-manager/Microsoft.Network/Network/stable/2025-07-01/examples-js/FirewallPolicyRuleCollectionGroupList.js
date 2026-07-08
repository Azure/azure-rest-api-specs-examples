const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
 *
 * @summary lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
 * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupList.json
 */
async function listAllFirewallPolicyRuleCollectionGroupsForAGivenFirewallPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.firewallPolicyRuleCollectionGroups.list(
    "rg1",
    "firewallPolicy",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
