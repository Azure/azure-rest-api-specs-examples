const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a database's blob auditing policy.
 *
 * @summary creates or updates a database's blob auditing policy.
 * x-ms-original-file: 2025-01-01/DatabaseAzureMonitorAuditingCreateMin.json
 */
async function createOrUpdateADatabaseAzureMonitorAuditingPolicyWithMinimalParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseBlobAuditingPolicies.createOrUpdate(
    "blobauditingtest-4799",
    "blobauditingtest-6440",
    "testdb",
    { isAzureMonitorTargetEnabled: true, state: "Enabled" },
  );
  console.log(result);
}
