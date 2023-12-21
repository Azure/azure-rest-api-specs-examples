const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an extended database's blob auditing policy.
 *
 * @summary Creates or updates an extended database's blob auditing policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ExtendedDatabaseAzureMonitorAuditingCreateMin.json
 */
async function createOrUpdateAnExtendedDatabaseAzureMonitorAuditingPolicyWithMinimalParameters() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "blobauditingtest-4799";
  const serverName = "blobauditingtest-6440";
  const databaseName = "testdb";
  const parameters = {
    isAzureMonitorTargetEnabled: true,
    state: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.extendedDatabaseBlobAuditingPolicies.createOrUpdate(
    resourceGroupName,
    serverName,
    databaseName,
    parameters,
  );
  console.log(result);
}
