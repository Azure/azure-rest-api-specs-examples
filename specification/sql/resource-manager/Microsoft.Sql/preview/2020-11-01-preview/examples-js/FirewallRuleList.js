const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of firewall rules.
 *
 * @summary Gets a list of firewall rules.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FirewallRuleList.json
 */
async function listFirewallRules() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "firewallrulecrudtest-12";
  const serverName = "firewallrulecrudtest-6285";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.firewallRules.listByServer(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listFirewallRules().catch(console.error);
