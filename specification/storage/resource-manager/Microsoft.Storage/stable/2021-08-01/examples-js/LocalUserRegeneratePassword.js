const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

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
