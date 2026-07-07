const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithHttpHeadersToInsert.json
 */
async function createFirewallPolicyRuleCollectionGroupWithHttpHeaderToInsert() {
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
          action: { type: "Allow" },
          ruleCollectionType: "FirewallPolicyFilterRuleCollection",
          rules: [
            {
              name: "rule1",
              description: "Insert trusted tenants header",
              fqdnTags: ["WindowsVirtualDesktop"],
              httpHeadersToInsert: [
                {
                  headerName: "Restrict-Access-To-Tenants",
                  headerValue: "contoso.com,fabrikam.onmicrosoft.com",
                },
              ],
              protocols: [{ port: 80, protocolType: "Http" }],
              ruleType: "ApplicationRule",
              sourceAddresses: ["216.58.216.164", "10.0.0.0/24"],
            },
          ],
        },
      ],
    },
  );
  console.log(result);
}
