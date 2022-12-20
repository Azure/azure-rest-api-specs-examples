const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a storage account in Microsoft Azure.
 *
 * @summary Deletes a storage account in Microsoft Azure.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountDelete.json
 */
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
