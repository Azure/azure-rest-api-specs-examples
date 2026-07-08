const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithWebCategoriesPut.json
 */
async function createFirewallPolicyRuleCollectionGroupWithWebCategories() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "e747cc13-97d4-4a79-b463-42d7f4e558f2";
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
    },
  );
  console.log(result);
}
