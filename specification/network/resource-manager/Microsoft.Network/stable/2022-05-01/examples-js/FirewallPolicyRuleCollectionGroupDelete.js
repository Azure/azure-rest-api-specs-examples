const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary Deletes the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyRuleCollectionGroupDelete.json
 */
async function deleteFirewallPolicyRuleCollectionGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const ruleCollectionGroupName = "ruleCollectionGroup1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyRuleCollectionGroups.beginDeleteAndWait(
    resourceGroupName,
    firewallPolicyName,
    ruleCollectionGroupName
  );
  console.log(result);
}

deleteFirewallPolicyRuleCollectionGroup().catch(console.error);
