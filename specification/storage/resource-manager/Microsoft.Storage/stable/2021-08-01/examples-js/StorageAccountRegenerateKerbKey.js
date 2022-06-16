const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountRegenerateKerbKey() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4167";
  const accountName = "sto3539";
  const regenerateKey = {
    keyName: "kerb1",
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

storageAccountRegenerateKerbKey().catch(console.error);
