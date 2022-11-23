const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the private link resources that need to be created for a resource.
 *
 * @summary Get the private link resources that need to be created for a resource.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubPrivateLinkResources_List.json
 */
async function webPubSubPrivateLinkResourcesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webPubSubPrivateLinkResources.list(
    resourceGroupName,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

webPubSubPrivateLinkResourcesList().catch(console.error);
