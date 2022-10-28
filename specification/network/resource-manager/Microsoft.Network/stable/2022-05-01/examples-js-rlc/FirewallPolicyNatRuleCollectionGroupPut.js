const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyNatRuleCollectionGroupPut.json
 */
async function createFirewallPolicyNatRuleCollectionGroup() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const ruleCollectionGroupName = "ruleCollectionGroup1";
  const options = {
    body: {
      properties: {
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
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}",
      subscriptionId,
      resourceGroupName,
      firewallPolicyName,
      ruleCollectionGroupName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createFirewallPolicyNatRuleCollectionGroup().catch(console.error);
