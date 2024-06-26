const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the statistics from an IoT hub.
 *
 * @summary Get the statistics from an IoT hub.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/examples/iothub_stats.json
 */
async function iotHubResourceGetStats() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testHub";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.getStats(resourceGroupName, resourceName);
  console.log(result);
}

iotHubResourceGetStats().catch(console.error);
