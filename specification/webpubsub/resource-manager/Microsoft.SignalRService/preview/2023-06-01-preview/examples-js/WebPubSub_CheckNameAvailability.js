const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the resource name is valid and is not already in use.
 *
 * @summary Checks that the resource name is valid and is not already in use.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSub_CheckNameAvailability.json
 */
async function webPubSubCheckNameAvailability() {
  const subscriptionId =
    process.env["WEB-PUBSUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const parameters = {
    name: "myWebPubSubService",
    type: "Microsoft.SignalRService/WebPubSub",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSub.checkNameAvailability(location, parameters);
  console.log(result);
}
