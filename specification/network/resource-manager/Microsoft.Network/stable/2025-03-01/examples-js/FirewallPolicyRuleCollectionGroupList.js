const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
 *
 * @summary Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/FirewallPolicyRuleCollectionGroupList.json
 */
async function listAllFirewallPolicyRuleCollectionGroupsForAGivenFirewallPolicy() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const firewallPolicyName = "firewallPolicy";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.firewallPolicyRuleCollectionGroups.list(
    resourceGroupName,
    firewallPolicyName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
