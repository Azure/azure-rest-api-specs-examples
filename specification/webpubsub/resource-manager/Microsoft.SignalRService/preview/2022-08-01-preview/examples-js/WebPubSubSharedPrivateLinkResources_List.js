const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List shared private link resources
 *
 * @summary List shared private link resources
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubSharedPrivateLinkResources_List.json
 */
async function webPubSubSharedPrivateLinkResourcesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webPubSubSharedPrivateLinkResources.list(
    resourceGroupName,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

webPubSubSharedPrivateLinkResourcesList().catch(console.error);
