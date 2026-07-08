const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupPut.json
 */
async function createFirewallPolicyRuleCollectionGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyRuleCollectionGroups.createOrUpdate(
    "rg1",
    "firewallPolicy",
    "ruleCollectionGroup1",
    {
      priority: 100,
      ruleCollections: [
        {
          name: "Example-Filter-Rule-Collection",
          action: { type: "Deny" },
          priority: 100,
          ruleCollectionType: "FirewallPolicyFilterRuleCollection",
          rules: [
            {
              name: "network-rule1",
              destinationAddresses: ["*"],
              destinationPorts: ["*"],
              ipProtocols: ["TCP"],
              ruleType: "NetworkRule",
              sourceAddresses: ["10.1.25.0/24"],
            },
          ],
        },
      ],
    },
  );
  console.log(result);
}
