const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to restart a resource.
 *
 * @summary Operation to restart a resource.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSub_Restart.json
 */
async function webPubSubRestart() {
  const subscriptionId =
    process.env["WEB-PUBSUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WEB-PUBSUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSub.beginRestartAndWait(resourceGroupName, resourceName);
  console.log(result);
}
