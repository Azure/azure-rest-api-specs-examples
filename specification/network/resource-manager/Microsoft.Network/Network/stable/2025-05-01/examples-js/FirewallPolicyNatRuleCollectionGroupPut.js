const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/Network/stable/2025-05-01/examples/FirewallPolicyNatRuleCollectionGroupPut.json
 */
async function createFirewallPolicyNatRuleCollectionGroup() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
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
    parameters,
  );
  console.log(result);
}
