const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate the local user SSH password.
 *
 * @summary Regenerate the local user SSH password.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/LocalUserRegeneratePassword.json
 */
async function regenerateLocalUserPassword() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const username = "user1";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.localUsersOperations.regeneratePassword(
    resourceGroupName,
    accountName,
    username
  );
  console.log(result);
}

regenerateLocalUserPassword().catch(console.error);
