const { IotHubClient } = require("@azure/arm-iothub-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Test the new route for this Iot Hub
 *
 * @summary Test the new route for this Iot Hub
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2019-07-01-preview/examples/iothub_testnewroute.json
 */
async function iotHubResourceTestRoute() {
  const subscriptionId =
    process.env["IOTHUB_SUBSCRIPTION_ID"] || "91d12660-3dec-467a-be2a-213b5544ddc0";
  const iotHubName = "testHub";
  const resourceGroupName = process.env["IOTHUB_RESOURCE_GROUP"] || "myResourceGroup";
  const input = {
    message: {
      appProperties: {},
      body: "Body of message",
      systemProperties: {},
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
