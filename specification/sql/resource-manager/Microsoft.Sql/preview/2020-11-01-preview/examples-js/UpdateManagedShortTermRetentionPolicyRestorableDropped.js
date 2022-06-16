const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets a database's short term retention policy.
 *
 * @summary Sets a database's short term retention policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/UpdateManagedShortTermRetentionPolicyRestorableDropped.json
 */
async function updateTheShortTermRetentionPolicyForTheRestorableDroppedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup";
  const managedInstanceName = "testsvr";
  const restorableDroppedDatabaseId = "testdb,131403269876900000";
  const policyName = "default";
  const parameters = {
    retentionDays: 14,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result =
    await client.managedRestorableDroppedDatabaseBackupShortTermRetentionPolicies.beginUpdateAndWait(
      resourceGroupName,
      managedInstanceName,
      restorableDroppedDatabaseId,
      policyName,
      parameters
    );
  console.log(result);
}

updateTheShortTermRetentionPolicyForTheRestorableDroppedDatabase().catch(console.error);
