const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns communication details for a support ticket.
 *
 * @summary Returns communication details for a support ticket.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/GetCommunicationDetailsForSubscriptionSupportTicket.json
 */
async function getCommunicationDetailsForASubscriptionSupportTicket() {
  const subscriptionId = process.env["SUPPORT_SUBSCRIPTION_ID"] || "subid";
  const supportTicketName = "testticket";
  const communicationName = "testmessage";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.communications.get(supportTicketName, communicationName);
  console.log(result);
}
