const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of specified private endpoint connection
 *
 * @summary Update the state of specified private endpoint connection
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubPrivateEndpointConnections_Update.json
 */
async function webPubSubPrivateEndpointConnectionsUpdate() {
  const subscriptionId =
    process.env["WEB-PUBSUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const privateEndpointConnectionName = "mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e";
  const resourceGroupName = process.env["WEB-PUBSUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const parameters = {
    privateEndpoint: {},
    privateLinkServiceConnectionState: {
      actionsRequired: "None",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSubPrivateEndpointConnections.update(
    privateEndpointConnectionName,
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}
