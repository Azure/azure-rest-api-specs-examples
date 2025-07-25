const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a list of all verified partners.
 *
 * @summary Get a list of all verified partners.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/VerifiedPartners_List.json
 */
async function verifiedPartnersList() {
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.verifiedPartners.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
