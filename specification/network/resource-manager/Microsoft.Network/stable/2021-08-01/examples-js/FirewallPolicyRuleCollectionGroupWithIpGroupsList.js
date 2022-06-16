const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
 *
 * @summary Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/FirewallPolicyRuleCollectionGroupWithIpGroupsList.json
 */
async function listAllFirewallPolicyRuleCollectionGroupsWithIPGroupsForAGivenFirewallPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.firewallPolicyRuleCollectionGroups.list(
    resourceGroupName,
    firewallPolicyName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllFirewallPolicyRuleCollectionGroupsWithIPGroupsForAGivenFirewallPolicy().catch(console.error);
