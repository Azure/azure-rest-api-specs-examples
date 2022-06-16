const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/FirewallPolicyRuleCollectionGroupWithWebCategoriesPut.json
 */
async function createFirewallPolicyRuleCollectionGroupWithWebCategories() {
  const subscriptionId = "e747cc13-97d4-4a79-b463-42d7f4e558f2";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const ruleCollectionGroupName = "ruleCollectionGroup1";
  const parameters = {
    priority: 110,
    ruleCollections: [
      {
        name: "Example-Filter-Rule-Collection",
        action: { type: "Deny" },
        ruleCollectionType: "FirewallPolicyFilterRuleCollection",
        rules: [
          {
            name: "rule1",
            description: "Deny inbound rule",
            protocols: [{ port: 443, protocolType: "Https" }],
            ruleType: "ApplicationRule",
            sourceAddresses: ["216.58.216.164", "10.0.0.0/24"],
            webCategories: ["Hacking"],
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

createFirewallPolicyRuleCollectionGroupWithWebCategories().catch(console.error);
