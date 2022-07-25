const { DeviceUpdate } = require("@azure/arm-deviceupdate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of specified private endpoint connection associated with the device update account.
 *
 * @summary Update the state of specified private endpoint connection associated with the device update account.
 * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/PrivateEndpointConnections/PrivateEndpointConnection_CreateOrUpdate.json
 */
async function privateEndpointConnectionCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "test-rg";
  const accountName = "contoso";
  const privateEndpointConnectionName = "peexample01";
  const privateEndpointConnection = {
    privateLinkServiceConnectionState: {
      description: "Auto-Approved",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DeviceUpdate(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    privateEndpointConnectionName,
    privateEndpointConnection
  );
  console.log(result);
}

privateEndpointConnectionCreateOrUpdate().catch(console.error);
