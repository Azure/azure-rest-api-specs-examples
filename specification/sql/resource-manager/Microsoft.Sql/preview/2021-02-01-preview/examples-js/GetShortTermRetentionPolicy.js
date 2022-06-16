const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a database's short term retention policy.
 *
 * @summary Gets a database's short term retention policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/GetShortTermRetentionPolicy.json
 */
async function getTheShortTermRetentionPolicyForTheDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const policyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.backupShortTermRetentionPolicies.get(
    resourceGroupName,
    serverName,
    databaseName,
    policyName
  );
  console.log(result);
}

getTheShortTermRetentionPolicyForTheDatabase().catch(console.error);
