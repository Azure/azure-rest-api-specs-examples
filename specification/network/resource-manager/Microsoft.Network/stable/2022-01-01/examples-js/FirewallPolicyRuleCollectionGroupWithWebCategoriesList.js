const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
 *
 * @summary Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/FirewallPolicyRuleCollectionGroupWithWebCategoriesList.json
 */
async function listAllFirewallPolicyRuleCollectionGroupWithWebCategories() {
  const subscriptionId = "e747cc13-97d4-4a79-b463-42d7f4e558f2";
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

listAllFirewallPolicyRuleCollectionGroupWithWebCategories().catch(console.error);
