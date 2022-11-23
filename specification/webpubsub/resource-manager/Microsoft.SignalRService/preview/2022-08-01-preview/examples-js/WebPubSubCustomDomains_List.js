const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all custom domains.
 *
 * @summary List all custom domains.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubCustomDomains_List.json
 */
async function webPubSubCustomDomainsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webPubSubCustomDomains.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

webPubSubCustomDomainsList().catch(console.error);
