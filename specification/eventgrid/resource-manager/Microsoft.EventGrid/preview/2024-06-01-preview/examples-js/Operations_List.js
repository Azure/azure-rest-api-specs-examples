const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the available operations supported by the Microsoft.EventGrid resource provider.
 *
 * @summary List the available operations supported by the Microsoft.EventGrid resource provider.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/Operations_List.json
 */
async function operationsList() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
