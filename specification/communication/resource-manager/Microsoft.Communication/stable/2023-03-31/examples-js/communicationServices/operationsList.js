const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available REST API operations of the Microsoft.Communication provider.
 *
 * @summary Lists all of the available REST API operations of the Microsoft.Communication provider.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/communicationServices/operationsList.json
 */
async function operationsList() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
