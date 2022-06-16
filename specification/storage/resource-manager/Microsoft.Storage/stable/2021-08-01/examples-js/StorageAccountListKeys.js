const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountListKeys() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res418";
  const accountName = "sto2220";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.listKeys(resourceGroupName, accountName);
  console.log(result);
}

storageAccountListKeys().catch(console.error);
