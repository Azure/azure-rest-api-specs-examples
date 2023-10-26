const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List hub settings.
 *
 * @summary List hub settings.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubHubs_List.json
 */
async function webPubSubHubsList() {
  const subscriptionId =
    process.env["WEB-PUBSUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WEB-PUBSUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webPubSubHubs.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
