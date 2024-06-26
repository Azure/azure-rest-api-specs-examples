const { IotCentralClient } = require("@azure/arm-iotcentral");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a private endpoint connection from the IoT Central Application.
 *
 * @summary Deletes a private endpoint connection from the IoT Central Application.
 * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_Delete.json
 */
async function privateEndpointConnectionsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const resourceName = "myIoTCentralApp";
  const privateEndpointConnectionName = "myIoTCentralAppEndpoint";
  const credential = new DefaultAzureCredential();
  const client = new IotCentralClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginDeleteAndWait(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName
  );
  console.log(result);
}

privateEndpointConnectionsDelete().catch(console.error);
