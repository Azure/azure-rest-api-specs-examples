const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a firewall rule.
 *
 * @summary Deletes a firewall rule.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/Firewall/preview/2023-06-01-preview/examples/FirewallRuleDelete.json
 */
async function deleteAFirewallRule() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["MYSQL_RESOURCE_GROUP"] || "TestGroup";
  const serverName = "testserver";
  const firewallRuleName = "rule1";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.firewallRules.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    firewallRuleName,
  );
  console.log(result);
}
