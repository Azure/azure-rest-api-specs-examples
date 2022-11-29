const { ManagementLockClient } = require("@azure/arm-locks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the management locks for a subscription.
 *
 * @summary Gets all the management locks for a subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_ListAtSubscriptionLevel.json
 */
async function listManagementLocksAtSubscriptionLevel() {
  const subscriptionId = "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new ManagementLockClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managementLocks.listAtSubscriptionLevel()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagementLocksAtSubscriptionLevel().catch(console.error);
