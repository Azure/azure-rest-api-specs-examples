const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a firewall rule.
 *
 * @summary Gets a firewall rule.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FirewallRuleGet.json
 */
async function getFirewallRule() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "firewallrulecrudtest-12";
  const serverName = "firewallrulecrudtest-6285";
  const firewallRuleName = "firewallrulecrudtest-2304";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.get(resourceGroupName, serverName, firewallRuleName);
  console.log(result);
}

getFirewallRule().catch(console.error);
