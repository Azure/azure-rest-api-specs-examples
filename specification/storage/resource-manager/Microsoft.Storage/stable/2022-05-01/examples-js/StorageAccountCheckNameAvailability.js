const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the storage account name is valid and is not already in use.
 *
 * @summary Checks that the storage account name is valid and is not already in use.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountCheckNameAvailability.json
 */
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
