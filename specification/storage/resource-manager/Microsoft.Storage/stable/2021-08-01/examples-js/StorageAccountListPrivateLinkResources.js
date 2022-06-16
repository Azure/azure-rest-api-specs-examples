const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountListPrivateLinkResources() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.listByStorageAccount(
    resourceGroupName,
    accountName
  );
  console.log(result);
}

storageAccountListPrivateLinkResources().catch(console.error);
