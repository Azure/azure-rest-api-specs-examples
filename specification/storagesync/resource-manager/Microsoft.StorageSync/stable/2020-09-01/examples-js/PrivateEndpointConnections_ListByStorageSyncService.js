const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a PrivateEndpointConnection List.
 *
 * @summary Get a PrivateEndpointConnection List.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/PrivateEndpointConnections_ListByStorageSyncService.json
 */
async function privateEndpointConnectionsListByStorageSyncService() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const storageSyncServiceName = "sss2527";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.listByStorageSyncService(
    resourceGroupName,
    storageSyncServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateEndpointConnectionsListByStorageSyncService().catch(console.error);
