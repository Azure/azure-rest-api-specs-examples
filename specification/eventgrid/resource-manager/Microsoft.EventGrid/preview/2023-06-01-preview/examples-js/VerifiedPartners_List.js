const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all verified partners.
 *
 * @summary Get a list of all verified partners.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/VerifiedPartners_List.json
 */
async function verifiedPartnersList() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.verifiedPartners.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
