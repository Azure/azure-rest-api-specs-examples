const { StorageActionsManagementClient } = require("@azure/arm-storageactions");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists Resource IDs of the Storage Task Assignments associated with this Storage Task.
 *
 * @summary Lists Resource IDs of the Storage Task Assignments associated with this Storage Task.
 * x-ms-original-file: specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/storageTasksList/ListStorageTaskAssignmentIds.json
 */
async function listStorageTaskAssignmentsByResourceGroup() {
  const subscriptionId =
    process.env["STORAGEACTIONS_SUBSCRIPTION_ID"] || "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
  const resourceGroupName = process.env["STORAGEACTIONS_RESOURCE_GROUP"] || "rgroup1";
  const storageTaskName = "mytask1";
  const credential = new DefaultAzureCredential();
  const client = new StorageActionsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.storageTaskAssignmentOperations.list(
    resourceGroupName,
    storageTaskName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
