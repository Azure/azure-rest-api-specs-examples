const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyRuleCollectionGroupWithIpGroupsPut.json
 */
async function createFirewallPolicyRuleCollectionGroupWithIPGroups() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const ruleCollectionGroupName = "ruleCollectionGroup1";
  const options = {
    body: {
      properties: {
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
                  "/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups2",
                ],
                destinationPorts: ["*"],
                ipProtocols: ["TCP"],
                ruleType: "NetworkRule",
                sourceIpGroups: [
                  "/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups1",
                ],
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

createFirewallPolicyRuleCollectionGroupWithIPGroups().catch(console.error);
