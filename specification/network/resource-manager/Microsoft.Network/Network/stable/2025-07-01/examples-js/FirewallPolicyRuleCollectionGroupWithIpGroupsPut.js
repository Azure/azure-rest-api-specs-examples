const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithIpGroupsPut.json
 */
async function createFirewallPolicyRuleCollectionGroupWithIPGroups() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyRuleCollectionGroups.createOrUpdate(
    "rg1",
    "firewallPolicy",
    "ruleCollectionGroup1",
    {
      priority: 110,
      ruleCollections: [
        {
          name: "Example-Filter-Rule-Collection",
          action: { type: "Deny" },
          ruleCollectionType: "FirewallPolicyFilterRuleCollection",
          rules: [
            {
              name: "network-1",
              destinationIpGroups: [
                "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups2",
              ],
              destinationPorts: ["*"],
              ipProtocols: ["TCP"],
              ruleType: "NetworkRule",
              sourceIpGroups: [
                "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/rg1/ipGroups/ipGroups1",
              ],
            },
          ],
        },
      ],
    },
  );
  console.log(result);
}
