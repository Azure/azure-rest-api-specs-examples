const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to returns the properties for the specified storage account including but not limited to name, SKU name, location, and account status. The ListKeys operation should be used to retrieve storage keys.
 *
 * @summary returns the properties for the specified storage account including but not limited to name, SKU name, location, and account status. The ListKeys operation should be used to retrieve storage keys.
 * x-ms-original-file: 2026-04-01/StorageAccountGetPropertiesCMKEnabled.json
 */
async function storageAccountGetPropertiesCMKEnabled() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.getProperties("res9407", "sto8596");
  console.log(result);
}
