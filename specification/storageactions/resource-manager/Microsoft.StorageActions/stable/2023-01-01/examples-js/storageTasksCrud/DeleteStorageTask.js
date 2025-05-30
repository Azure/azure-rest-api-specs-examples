const { StorageActionsManagementClient } = require("@azure/arm-storageactions");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete the storage task resource.
 *
 * @summary Delete the storage task resource.
 * x-ms-original-file: specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/storageTasksCrud/DeleteStorageTask.json
 */
async function deleteStorageTask() {
  const subscriptionId =
    process.env["STORAGEACTIONS_SUBSCRIPTION_ID"] || "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
  const resourceGroupName = process.env["STORAGEACTIONS_RESOURCE_GROUP"] || "res4228";
  const storageTaskName = "mytask1";
  const credential = new DefaultAzureCredential();
  const client = new StorageActionsManagementClient(credential, subscriptionId);
  const result = await client.storageTasks.beginDeleteAndWait(resourceGroupName, storageTaskName);
  console.log(result);
}
