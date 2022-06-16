const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountGetProperties() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9407";
  const accountName = "sto8596";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.getProperties(resourceGroupName, accountName);
  console.log(result);
}

storageAccountGetProperties().catch(console.error);
