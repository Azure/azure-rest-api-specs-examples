const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified shared private link resource
 *
 * @summary Get the specified shared private link resource
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/WebPubSubReplicaSharedPrivateLinkResources_Get.json
 */
async function webPubSubReplicaSharedPrivateLinkResourcesGet() {
  const subscriptionId =
    process.env["WEB-PUBSUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WEB-PUBSUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const replicaName = "myWebPubSubService-eastus";
  const sharedPrivateLinkResourceName = "upstream";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSubReplicaSharedPrivateLinkResources.get(
    resourceGroupName,
    resourceName,
    replicaName,
    sharedPrivateLinkResourceName,
  );
  console.log(result);
}
