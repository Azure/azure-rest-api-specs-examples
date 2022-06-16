const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets properties of a specified share.
 *
 * @summary Gets properties of a specified share.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/FileSharesGet.json
 */
async function getShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9871";
  const accountName = "sto6217";
  const shareName = "share1634";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.get(resourceGroupName, accountName, shareName);
  console.log(result);
}

getShares().catch(console.error);
