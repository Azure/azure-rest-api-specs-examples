const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of specified private endpoint connection associated with the storage sync service.
 *
 * @summary Update the state of specified private endpoint connection associated with the storage sync service.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/PrivateEndpointConnections_Create.json
 */
async function privateEndpointConnectionsCreate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const storageSyncServiceName = "sss2527";
  const privateEndpointConnectionName = "{privateEndpointConnectionName}";
  const properties = {
    privateLinkServiceConnectionState: {
      description: "Auto-Approved",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateAndWait(
    resourceGroupName,
    storageSyncServiceName,
    privateEndpointConnectionName,
    properties
  );
  console.log(result);
}

privateEndpointConnectionsCreate().catch(console.error);
