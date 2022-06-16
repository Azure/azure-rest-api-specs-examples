const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a dropped database's short term retention policy.
 *
 * @summary Gets a dropped database's short term retention policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetManagedShortTermRetentionPolicyRestorableDropped.json
 */
async function getTheShortTermRetentionPolicyForTheDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const managedInstanceName = "testsvr";
  const restorableDroppedDatabaseId = "testdb,131403269876900000";
  const policyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedRestorableDroppedDatabaseBackupShortTermRetentionPolicies.get(
    resourceGroupName,
    managedInstanceName,
    restorableDroppedDatabaseId,
    policyName
  );
  console.log(result);
}

getTheShortTermRetentionPolicyForTheDatabase().catch(console.error);
