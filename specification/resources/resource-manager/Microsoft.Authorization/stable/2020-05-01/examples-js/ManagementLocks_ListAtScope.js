const { ManagementLockClient } = require("@azure/arm-locks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the management locks for a scope.
 *
 * @summary Gets all the management locks for a scope.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_ListAtScope.json
 */
async function listManagementLocksAtScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new ManagementLockClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managementLocks.listByScope(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagementLocksAtScope().catch(console.error);
