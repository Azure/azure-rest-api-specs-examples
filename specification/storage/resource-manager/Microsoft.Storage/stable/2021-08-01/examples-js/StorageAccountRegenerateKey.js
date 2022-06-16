const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountRegenerateKey() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4167";
  const accountName = "sto3539";
  const regenerateKey = {
    keyName: "key2",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.regenerateKey(
    resourceGroupName,
    accountName,
    regenerateKey
  );
  console.log(result);
}

storageAccountRegenerateKey().catch(console.error);
