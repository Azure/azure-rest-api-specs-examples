const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Storage Movers in a resource group.
 *
 * @summary Lists all Storage Movers in a resource group.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/StorageMovers_List.json
 */
async function storageMoversList() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.storageMovers.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
