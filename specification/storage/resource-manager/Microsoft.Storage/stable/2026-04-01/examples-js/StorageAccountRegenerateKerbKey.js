const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to regenerates one of the access keys or Kerberos keys for the specified storage account.
 *
 * @summary regenerates one of the access keys or Kerberos keys for the specified storage account.
 * x-ms-original-file: 2026-04-01/StorageAccountRegenerateKerbKey.json
 */
async function storageAccountRegenerateKerbKey() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.regenerateKey("res4167", "sto3539", {
    keyName: "kerb1",
  });
  console.log(result);
}
