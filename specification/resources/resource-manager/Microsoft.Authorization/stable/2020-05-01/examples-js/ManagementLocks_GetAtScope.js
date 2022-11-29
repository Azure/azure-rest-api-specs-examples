const { ManagementLockClient } = require("@azure/arm-locks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a management lock by scope.
 *
 * @summary Get a management lock by scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_GetAtScope.json
 */
async function getManagementLockAtScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/subscriptionId";
  const lockName = "testlock";
  const credential = new DefaultAzureCredential();
  const client = new ManagementLockClient(credential, subscriptionId);
  const result = await client.managementLocks.getByScope(scope, lockName);
  console.log(result);
}

getManagementLockAtScope().catch(console.error);
