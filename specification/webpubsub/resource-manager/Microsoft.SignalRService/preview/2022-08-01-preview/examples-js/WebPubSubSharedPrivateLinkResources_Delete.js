const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the specified shared private link resource
 *
 * @summary Delete the specified shared private link resource
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubSharedPrivateLinkResources_Delete.json
 */
async function webPubSubSharedPrivateLinkResourcesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const sharedPrivateLinkResourceName = "upstream";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSubSharedPrivateLinkResources.beginDeleteAndWait(
    sharedPrivateLinkResourceName,
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

webPubSubSharedPrivateLinkResourcesDelete().catch(console.error);
