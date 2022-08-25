const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all resources in a subscription.
 *
 * @summary Handles requests to list all resources in a subscription.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/emailServices/listBySubscription.json
 */
async function listEmailServiceResourcesBySubscription() {
  const subscriptionId = "12345";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.emailServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listEmailServiceResourcesBySubscription().catch(console.error);
