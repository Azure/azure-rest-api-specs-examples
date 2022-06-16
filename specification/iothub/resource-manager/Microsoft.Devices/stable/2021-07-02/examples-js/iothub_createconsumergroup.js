const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add a consumer group to an Event Hub-compatible endpoint in an IoT hub.
 *
 * @summary Add a consumer group to an Event Hub-compatible endpoint in an IoT hub.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_createconsumergroup.json
 */
async function iotHubResourceCreateEventHubConsumerGroup() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testHub";
  const eventHubEndpointName = "events";
  const name = "test";
  const consumerGroupBody = {
    properties: { name: "test" },
  };
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.createEventHubConsumerGroup(
    resourceGroupName,
    resourceName,
    eventHubEndpointName,
    name,
    consumerGroupBody
  );
  console.log(result);
}

iotHubResourceCreateEventHubConsumerGroup().catch(console.error);
