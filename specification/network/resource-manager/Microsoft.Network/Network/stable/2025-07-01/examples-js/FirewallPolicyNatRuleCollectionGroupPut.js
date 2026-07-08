const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: 2025-07-01/FirewallPolicyNatRuleCollectionGroupPut.json
 */
async function createFirewallPolicyNatRuleCollectionGroup() {
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
          name: "Example-Nat-Rule-Collection",
          action: { type: "DNAT" },
          priority: 100,
          ruleCollectionType: "FirewallPolicyNatRuleCollection",
          rules: [
            {
              name: "nat-rule1",
              destinationAddresses: ["152.23.32.23"],
              destinationPorts: ["8080"],
              ipProtocols: ["TCP", "UDP"],
              ruleType: "NatRule",
              sourceAddresses: ["2.2.2.2"],
              sourceIpGroups: [],
              translatedFqdn: "internalhttp.server.net",
              translatedPort: "8080",
            },
          ],
        },
      ],
    },
  );
  console.log(result);
}
