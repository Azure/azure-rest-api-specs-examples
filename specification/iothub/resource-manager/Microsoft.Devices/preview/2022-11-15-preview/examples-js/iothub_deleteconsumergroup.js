const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a consumer group from an Event Hub-compatible endpoint in an IoT hub.
 *
 * @summary Delete a consumer group from an Event Hub-compatible endpoint in an IoT hub.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_deleteconsumergroup.json
 */
async function iotHubResourceDeleteEventHubConsumerGroup() {
  const subscriptionId =
    process.env["IOTHUB_SUBSCRIPTION_ID"] || "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = process.env["IOTHUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "testHub";
  const eventHubEndpointName = "events";
  const name = "test";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.deleteEventHubConsumerGroup(
    resourceGroupName,
    resourceName,
    eventHubEndpointName,
    name
  );
  console.log(result);
}
