const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all resources in a subscription.
 *
 * @summary Handles requests to list all resources in a subscription.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/communicationServices/listBySubscription.json
 */
async function listBySubscription() {
  const subscriptionId = process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "12345";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communicationServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
