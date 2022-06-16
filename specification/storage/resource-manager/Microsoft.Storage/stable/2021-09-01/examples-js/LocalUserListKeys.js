const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List SSH authorized keys and shared key of the local user.
 *
 * @summary List SSH authorized keys and shared key of the local user.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/LocalUserListKeys.json
 */
async function listLocalUserKeys() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const username = "user1";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.localUsersOperations.listKeys(
    resourceGroupName,
    accountName,
    username
  );
  console.log(result);
}

listLocalUserKeys().catch(console.error);
