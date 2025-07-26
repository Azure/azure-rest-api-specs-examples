const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get properties of a verified partner.
 *
 * @summary Get properties of a verified partner.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/VerifiedPartners_Get.json
 */
async function verifiedPartnersGet() {
  const verifiedPartnerName = "Contoso.Finance";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential);
  const result = await client.verifiedPartners.get(verifiedPartnerName);
  console.log(result);
}
