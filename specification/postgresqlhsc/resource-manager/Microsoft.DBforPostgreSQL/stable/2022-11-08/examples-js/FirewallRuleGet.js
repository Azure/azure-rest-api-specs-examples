const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about a cluster firewall rule.
 *
 * @summary Gets information about a cluster firewall rule.
 * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/FirewallRuleGet.json
 */
async function getTheFirewallRuleOfTheCluster() {
  const subscriptionId =
    process.env["COSMOSFORPOSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSFORPOSTGRESQL_RESOURCE_GROUP"] || "TestGroup";
  const clusterName = "pgtestsvc4";
  const firewallRuleName = "rule1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.firewallRules.get(resourceGroupName, clusterName, firewallRuleName);
  console.log(result);
}
