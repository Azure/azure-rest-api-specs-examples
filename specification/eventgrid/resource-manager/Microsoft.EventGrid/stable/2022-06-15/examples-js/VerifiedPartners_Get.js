const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get properties of a verified partner.
 *
 * @summary Get properties of a verified partner.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/VerifiedPartners_Get.json
 */
async function verifiedPartnersGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const verifiedPartnerName = "Contoso.Finance";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.verifiedPartners.get(verifiedPartnerName);
  console.log(result);
}

verifiedPartnersGet().catch(console.error);
