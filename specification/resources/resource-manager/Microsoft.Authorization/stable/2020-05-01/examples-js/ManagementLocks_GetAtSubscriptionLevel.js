const { ManagementLockClient } = require("@azure/arm-locks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a management lock at the subscription level.
 *
 * @summary Gets a management lock at the subscription level.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_GetAtSubscriptionLevel.json
 */
async function getManagementLockAtSubscriptionLevel() {
  const subscriptionId = "subscriptionId";
  const lockName = "testlock";
  const credential = new DefaultAzureCredential();
  const client = new ManagementLockClient(credential, subscriptionId);
  const result = await client.managementLocks.getAtSubscriptionLevel(lockName);
  console.log(result);
}

getManagementLockAtSubscriptionLevel().catch(console.error);
