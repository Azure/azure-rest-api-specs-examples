const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a firewall rule.
 *
 * @summary creates or updates a firewall rule.
 * x-ms-original-file: 2025-01-01/FirewallRuleCreate.json
 */
async function createAFirewallRuleMaxOrMin() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.createOrUpdate(
    "firewallrulecrudtest-12",
    "firewallrulecrudtest-6285",
    "firewallrulecrudtest-5370",
    { endIpAddress: "0.0.0.3", startIpAddress: "0.0.0.3" },
  );
  console.log(result);
}
