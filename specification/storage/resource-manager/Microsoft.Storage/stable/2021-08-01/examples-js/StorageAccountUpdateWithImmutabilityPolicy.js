const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountUpdateWithImmutabilityPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9407";
  const accountName = "sto8596";
  const parameters = {
    immutableStorageWithVersioning: {
      enabled: true,
      immutabilityPolicy: {
        allowProtectedAppendWrites: true,
        immutabilityPeriodSinceCreationInDays: 15,
        state: "Locked",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.update(resourceGroupName, accountName, parameters);
  console.log(result);
}

storageAccountUpdateWithImmutabilityPolicy().catch(console.error);
