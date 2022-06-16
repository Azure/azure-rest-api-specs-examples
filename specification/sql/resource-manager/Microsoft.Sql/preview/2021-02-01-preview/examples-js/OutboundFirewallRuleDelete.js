const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a outbound firewall rule with a given name.
 *
 * @summary Deletes a outbound firewall rule with a given name.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/OutboundFirewallRuleDelete.json
 */
async function deletesAOutboundFirewallRuleWithAGivenName() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const serverName = "sqlcrudtest-6661";
  const outboundRuleFqdn = "server.database.windows.net";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.outboundFirewallRules.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    outboundRuleFqdn
  );
  console.log(result);
}

deletesAOutboundFirewallRuleWithAGivenName().catch(console.error);
