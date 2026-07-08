const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary gets the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithWebCategoriesGet.json
 */
async function getFirewallPolicyRuleCollectionGroupWithWebCategories() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "e747cc13-97d4-4a79-b463-42d7f4e558f2";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyRuleCollectionGroups.get(
    "rg1",
    "firewallPolicy",
    "ruleCollectionGroup1",
  );
  console.log(result);
}
