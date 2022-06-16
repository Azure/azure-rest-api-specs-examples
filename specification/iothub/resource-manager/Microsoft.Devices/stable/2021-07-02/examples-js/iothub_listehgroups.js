const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of the consumer groups in the Event Hub-compatible device-to-cloud endpoint in an IoT hub.
 *
 * @summary Get a list of the consumer groups in the Event Hub-compatible device-to-cloud endpoint in an IoT hub.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listehgroups.json
 */
async function iotHubResourceListEventHubConsumerGroups() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testHub";
  const eventHubEndpointName = "events";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotHubResource.listEventHubConsumerGroups(
    resourceGroupName,
    resourceName,
    eventHubEndpointName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

iotHubResourceListEventHubConsumerGroups().catch(console.error);
