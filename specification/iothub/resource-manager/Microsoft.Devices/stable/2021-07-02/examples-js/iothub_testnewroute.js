const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Test the new route for this Iot Hub
 *
 * @summary Test the new route for this Iot Hub
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_testnewroute.json
 */
async function iotHubResourceTestRoute() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const iotHubName = "testHub";
  const resourceGroupName = "myResourceGroup";
  const input = {
    message: {
      appProperties: { key1: "value1" },
      body: "Body of message",
      systemProperties: { key1: "value1" },
    },
    route: {
      name: "Routeid",
      endpointNames: ["id1"],
      isEnabled: true,
      source: "DeviceMessages",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.testRoute(iotHubName, resourceGroupName, input);
  console.log(result);
}

iotHubResourceTestRoute().catch(console.error);
