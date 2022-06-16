const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the status of a private endpoint connection with the specified name
 *
 * @summary Update the status of a private endpoint connection with the specified name
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_updateprivateendpointconnection.json
 */
async function privateEndpointConnectionUpdate() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testHub";
  const privateEndpointConnectionName = "myPrivateEndpointConnection";
  const privateEndpointConnection = {
    properties: {
      privateLinkServiceConnectionState: {
        description: "Approved by johndoe@contoso.com",
        status: "Approved",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginUpdateAndWait(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName,
    privateEndpointConnection
  );
  console.log(result);
}

privateEndpointConnectionUpdate().catch(console.error);
