const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all shares.
 *
 * @summary Lists all shares.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/FileSharesList.json
 */
async function listShares() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9290";
  const accountName = "sto1590";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fileShares.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listShares().catch(console.error);
