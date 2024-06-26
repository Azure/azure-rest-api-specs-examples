const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of operations supported by the import/export resource provider.
 *
 * @summary Returns the list of operations supported by the import/export resource provider.
 * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListOperations.json
 */
async function listAvailableOperations() {
  const subscriptionId =
    process.env["STORAGEIMPORTEXPORT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new StorageImportExport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
