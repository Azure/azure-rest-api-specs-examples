const { IotHubClient } = require("@azure/arm-iothub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the number of free and paid iot hubs in the subscription
 *
 * @summary Get the number of free and paid iot hubs in the subscription
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_usages.json
 */
async function resourceProviderCommonGetSubscriptionQuota() {
  const subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.resourceProviderCommon.getSubscriptionQuota();
  console.log(result);
}

resourceProviderCommonGetSubscriptionQuota().catch(console.error);
