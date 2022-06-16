const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all the IoT hubs in a subscription.
 *
 * @summary Get all the IoT hubs in a subscription.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listbysubscription.json
 */
async function iotHubResourceListBySubscription() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotHubResource.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

iotHubResourceListBySubscription().catch(console.error);
