const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all of the available Resource Provider operations.
 *
 * @summary lists all of the available Resource Provider operations.
 * x-ms-original-file: 2026-01-01/Operations_List.json
 */
async function operationsList() {
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.operations.list()) {
    resArray.push(item);
  }

  console.log(resArray);
}
