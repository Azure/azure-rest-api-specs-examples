const { StorageActionsManagementClient } = require("@azure/arm-storageactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all the storage tasks available under the given resource group.
 *
 * @summary lists all the storage tasks available under the given resource group.
 * x-ms-original-file: 2023-01-01/storageTasksList/ListStorageTasksByResourceGroup.json
 */
async function listStorageTasksByResourceGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
  const client = new StorageActionsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.storageTasks.listByResourceGroup("res6117")) {
    resArray.push(item);
  }

  console.log(resArray);
}
