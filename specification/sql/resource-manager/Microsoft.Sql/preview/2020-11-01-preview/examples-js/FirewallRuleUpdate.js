const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a firewall rule.
 *
 * @summary Creates or updates a firewall rule.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FirewallRuleUpdate.json
 */
async function updateAFirewallRuleMaxOrMin() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "firewallrulecrudtest-12";
  const serverName = "firewallrulecrudtest-6285";
  const firewallRuleName = "firewallrulecrudtest-3927";
  const parameters = {
    endIpAddress: "0.0.0.1",
    startIpAddress: "0.0.0.1",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.createOrUpdate(
    resourceGroupName,
    serverName,
    firewallRuleName,
    parameters
  );
  console.log(result);
}

updateAFirewallRuleMaxOrMin().catch(console.error);
