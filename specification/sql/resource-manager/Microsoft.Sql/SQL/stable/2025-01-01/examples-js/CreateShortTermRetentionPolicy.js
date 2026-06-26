const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a database's short term retention policy.
 *
 * @summary updates a database's short term retention policy.
 * x-ms-original-file: 2025-01-01/CreateShortTermRetentionPolicy.json
 */
async function updateTheShortTermRetentionPolicyForTheDatabase() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.backupShortTermRetentionPolicies.createOrUpdate(
    "resourceGroup",
    "testsvr",
    "testdb",
    "default",
    { diffBackupIntervalInHours: 24, retentionDays: 7 },
  );
  console.log(result);
}
