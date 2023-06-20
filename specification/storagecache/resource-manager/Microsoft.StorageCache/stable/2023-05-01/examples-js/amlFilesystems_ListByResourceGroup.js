const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all AML file systems the user has access to under a resource group.
 *
 * @summary Returns all AML file systems the user has access to under a resource group.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/examples/amlFilesystems_ListByResourceGroup.json
 */
async function amlFilesystemsListByResourceGroup() {
  const subscriptionId =
    process.env["STORAGECACHE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["STORAGECACHE_RESOURCE_GROUP"] || "scgroup";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.amlFilesystems.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
