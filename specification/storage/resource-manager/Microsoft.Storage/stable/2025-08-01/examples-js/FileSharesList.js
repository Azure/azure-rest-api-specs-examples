const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all shares.
 *
 * @summary lists all shares.
 * x-ms-original-file: 2025-08-01/FileSharesList.json
 */
async function listShares() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.fileShares.list("res9290", "sto1590")) {
    resArray.push(item);
  }

  console.log(resArray);
}
