const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountCheckNameAvailability() {
  const subscriptionId = "{subscription-id}";
  const accountName = {
    name: "sto3363",
    type: "Microsoft.Storage/storageAccounts",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.checkNameAvailability(accountName);
  console.log(result);
}

storageAccountCheckNameAvailability().catch(console.error);
