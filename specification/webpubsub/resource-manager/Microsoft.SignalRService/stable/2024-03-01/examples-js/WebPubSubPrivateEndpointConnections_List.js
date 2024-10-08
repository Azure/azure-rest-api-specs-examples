const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List private endpoint connections
 *
 * @summary List private endpoint connections
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/WebPubSubPrivateEndpointConnections_List.json
 */
async function webPubSubPrivateEndpointConnectionsList() {
  const subscriptionId =
    process.env["WEB-PUBSUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WEB-PUBSUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webPubSubPrivateEndpointConnections.list(
    resourceGroupName,
    resourceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
