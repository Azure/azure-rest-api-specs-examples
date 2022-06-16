const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/FirewallPolicyNatRuleCollectionGroupPut.json
 */
async function createFirewallPolicyNatRuleCollectionGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const ruleCollectionGroupName = "ruleCollectionGroup1";
  const parameters = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyRuleCollectionGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    firewallPolicyName,
    ruleCollectionGroupName,
    parameters
  );
  console.log(result);
}

createFirewallPolicyNatRuleCollectionGroup().catch(console.error);
