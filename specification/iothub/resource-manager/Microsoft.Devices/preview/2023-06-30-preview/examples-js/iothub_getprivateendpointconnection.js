const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get private endpoint connection properties
 *
 * @summary Get private endpoint connection properties
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_getprivateendpointconnection.json
 */
async function privateEndpointConnectionGet() {
  const subscriptionId =
    process.env["IOTHUB_SUBSCRIPTION_ID"] || "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = process.env["IOTHUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "testHub";
  const privateEndpointConnectionName = "myPrivateEndpointConnection";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName
  );
  console.log(result);
}
