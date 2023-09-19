const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a consumer group from the Event Hub-compatible device-to-cloud endpoint for an IoT hub.
 *
 * @summary Get a consumer group from the Event Hub-compatible device-to-cloud endpoint for an IoT hub.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_getconsumergroup.json
 */
async function iotHubResourceListEventHubConsumerGroups() {
  const subscriptionId =
    process.env["IOTHUB_SUBSCRIPTION_ID"] || "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = process.env["IOTHUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "testHub";
  const eventHubEndpointName = "events";
  const name = "test";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.getEventHubConsumerGroup(
    resourceGroupName,
    resourceName,
    eventHubEndpointName,
    name
  );
  console.log(result);
}
