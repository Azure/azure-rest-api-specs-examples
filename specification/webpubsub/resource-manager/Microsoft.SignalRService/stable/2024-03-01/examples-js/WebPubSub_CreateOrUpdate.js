const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a resource.
 *
 * @summary Create or update a resource.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/WebPubSub_CreateOrUpdate.json
 */
async function webPubSubCreateOrUpdate() {
  const subscriptionId =
    process.env["WEB-PUBSUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WEB-PUBSUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const parameters = {
    disableAadAuth: false,
    disableLocalAuth: false,
    identity: { type: "SystemAssigned" },
    kind: "WebPubSub",
    liveTraceConfiguration: {
      categories: [{ name: "ConnectivityLogs", enabled: "true" }],
      enabled: "false",
    },
    location: "eastus",
    networkACLs: {
      defaultAction: "Deny",
      privateEndpoints: [
        {
          name: "mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e",
          allow: ["ServerConnection"],
        },
      ],
      publicNetwork: { allow: ["ClientConnection"] },
    },
    publicNetworkAccess: "Enabled",
    sku: { name: "Premium_P1", capacity: 1, tier: "Premium" },
    socketIO: { serviceMode: "Serverless" },
    tags: { key1: "value1" },
    tls: { clientCertEnabled: false },
  };
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSub.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    parameters,
  );
  console.log(result);
}
