```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/FirewallPolicyRuleCollectionGroupPut.json
 */
async function createFirewallPolicyRuleCollectionGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const ruleCollectionGroupName = "ruleCollectionGroup1";
  const parameters = {
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

createFirewallPolicyRuleCollectionGroup().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
