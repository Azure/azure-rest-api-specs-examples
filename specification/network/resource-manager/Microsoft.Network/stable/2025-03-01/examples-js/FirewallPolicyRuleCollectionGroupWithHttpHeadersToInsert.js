const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 *
 * @summary Creates or updates the specified FirewallPolicyRuleCollectionGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/FirewallPolicyRuleCollectionGroupWithHttpHeadersToInsert.json
 */
async function createFirewallPolicyRuleCollectionGroupWithHttpHeaderToInsert() {
  const subscriptionId =
    process.env["NETWORK_SUBSCRIPTION_ID"] || "e747cc13-97d4-4a79-b463-42d7f4e558f2";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const firewallPolicyName = "firewallPolicy";
  const ruleCollectionGroupName = "ruleCollectionGroup1";
  const parameters = {
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
