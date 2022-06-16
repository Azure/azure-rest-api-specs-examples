const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the quota metrics for an IoT hub.
 *
 * @summary Get the quota metrics for an IoT hub.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_quotametrics.json
 */
async function iotHubResourceGetQuotaMetrics() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "testHub";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotHubResource.listQuotaMetrics(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

iotHubResourceGetQuotaMetrics().catch(console.error);
