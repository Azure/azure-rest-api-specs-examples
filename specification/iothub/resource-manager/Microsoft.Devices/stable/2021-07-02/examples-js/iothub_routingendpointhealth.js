const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the health for routing endpoints.
 *
 * @summary Get the health for routing endpoints.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_routingendpointhealth.json
 */
async function iotHubResourceGetEndpointHealth() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const iotHubName = "testHub";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotHubResource.listEndpointHealth(resourceGroupName, iotHubName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

iotHubResourceGetEndpointHealth().catch(console.error);
