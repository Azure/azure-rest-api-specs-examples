const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the properties for the specified storage account including but not limited to name, SKU name, location, and account status. The ListKeys operation should be used to retrieve storage keys.
 *
 * @summary Returns the properties for the specified storage account including but not limited to name, SKU name, location, and account status. The ListKeys operation should be used to retrieve storage keys.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountGetPropertiesCMKEnabled.json
 */
async function storageAccountGetPropertiesCmkEnabled() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9407";
  const accountName = "sto8596";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.getProperties(resourceGroupName, accountName);
  console.log(result);
}

storageAccountGetPropertiesCmkEnabled().catch(console.error);
