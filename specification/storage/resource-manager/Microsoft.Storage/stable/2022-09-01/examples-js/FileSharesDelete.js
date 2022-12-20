const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes specified share under its account.
 *
 * @summary Deletes specified share under its account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesDelete.json
 */
async function deleteShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4079";
  const accountName = "sto4506";
  const shareName = "share9689";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.delete(resourceGroupName, accountName, shareName);
  console.log(result);
}

deleteShares().catch(console.error);
