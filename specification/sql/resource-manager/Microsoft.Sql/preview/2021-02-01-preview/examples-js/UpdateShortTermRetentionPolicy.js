const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a database's short term retention policy.
 *
 * @summary Updates a database's short term retention policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/UpdateShortTermRetentionPolicy.json
 */
async function updateTheShortTermRetentionPolicyForTheDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const policyName = "default";
  const parameters = {
    diffBackupIntervalInHours: 24,
    retentionDays: 7,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.backupShortTermRetentionPolicies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    policyName,
    parameters
  );
  console.log(result);
}

updateTheShortTermRetentionPolicyForTheDatabase().catch(console.error);
