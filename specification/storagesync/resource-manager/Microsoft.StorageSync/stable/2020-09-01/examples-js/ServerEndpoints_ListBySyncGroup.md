Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagesync_9.0.1/sdk/storagesync/arm-storagesync/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a ServerEndpoint list.
 *
 * @summary Get a ServerEndpoint list.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_ListBySyncGroup.json
 */
async function serverEndpointsListBySyncGroup() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const syncGroupName = "SampleSyncGroup_1";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serverEndpoints.listBySyncGroup(
    resourceGroupName,
    storageSyncServiceName,
    syncGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

serverEndpointsListBySyncGroup().catch(console.error);
```
