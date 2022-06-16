const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List private endpoint connection properties
 *
 * @summary List private endpoint connection properties
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listprivateendpointconnections.json
 */
async function privateEndpointConnectionsList() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testHub";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.list(resourceGroupName, resourceName);
  console.log(result);
}

privateEndpointConnectionsList().catch(console.error);
