const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a managed database's long term retention policy.
 *
 * @summary Gets a managed database's long term retention policy.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceLongTermRetentionPolicyGet.json
 */
async function getTheLongTermRetentionPolicyForTheManagedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testResourceGroup";
  const managedInstanceName = "testInstance";
  const databaseName = "testDatabase";
  const policyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceLongTermRetentionPolicies.get(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    policyName
  );
  console.log(result);
}

getTheLongTermRetentionPolicyForTheManagedDatabase().catch(console.error);
