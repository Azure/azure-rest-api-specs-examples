const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all custom certificates.
 *
 * @summary List all custom certificates.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubCustomCertificates_List.json
 */
async function webPubSubCustomCertificatesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webPubSubCustomCertificates.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

webPubSubCustomCertificatesList().catch(console.error);
