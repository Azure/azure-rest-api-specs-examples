const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Storage Mover resource.
 *
 * @summary Gets a Storage Mover resource.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/StorageMovers_Get.json
 */
async function storageMoversGet() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.storageMovers.get(resourceGroupName, storageMoverName);
  console.log(result);
}
