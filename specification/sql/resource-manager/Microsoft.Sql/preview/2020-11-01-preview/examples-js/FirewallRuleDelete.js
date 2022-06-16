const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a firewall rule.
 *
 * @summary Deletes a firewall rule.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FirewallRuleDelete.json
 */
async function deleteAFirewallRule() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "firewallrulecrudtest-9886";
  const serverName = "firewallrulecrudtest-2368";
  const firewallRuleName = "firewallrulecrudtest-7011";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.delete(resourceGroupName, serverName, firewallRuleName);
  console.log(result);
}

deleteAFirewallRule().catch(console.error);
