const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all Caches the user has access to under a resource group.
 *
 * @summary Returns all Caches the user has access to under a resource group.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/Caches_ListByResourceGroup.json
 */
async function cachesListByResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.caches.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

cachesListByResourceGroup().catch(console.error);
