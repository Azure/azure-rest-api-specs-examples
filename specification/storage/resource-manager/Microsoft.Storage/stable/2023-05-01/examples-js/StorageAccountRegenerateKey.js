const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates one of the access keys or Kerberos keys for the specified storage account.
 *
 * @summary Regenerates one of the access keys or Kerberos keys for the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountRegenerateKey.json
 */
async function storageAccountRegenerateKey() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res4167";
  const accountName = "sto3539";
  const regenerateKey = {
    keyName: "key2",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.regenerateKey(
    resourceGroupName,
    accountName,
    regenerateKey,
  );
  console.log(result);
}
