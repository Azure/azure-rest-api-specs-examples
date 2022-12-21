const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates one of the access keys or Kerberos keys for the specified storage account.
 *
 * @summary Regenerates one of the access keys or Kerberos keys for the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountRegenerateKerbKey.json
 */
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
