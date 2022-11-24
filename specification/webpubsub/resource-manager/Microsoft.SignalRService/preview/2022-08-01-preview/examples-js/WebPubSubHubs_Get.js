const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a hub setting.
 *
 * @summary Get a hub setting.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubHubs_Get.json
 */
async function webPubSubHubsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const hubName = "exampleHub";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSubHubs.get(hubName, resourceGroupName, resourceName);
  console.log(result);
}

webPubSubHubsGet().catch(console.error);
