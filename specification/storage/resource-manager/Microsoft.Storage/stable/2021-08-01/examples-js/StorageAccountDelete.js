const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountDelete() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4228";
  const accountName = "sto2434";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.delete(resourceGroupName, accountName);
  console.log(result);
}

storageAccountDelete().catch(console.error);
