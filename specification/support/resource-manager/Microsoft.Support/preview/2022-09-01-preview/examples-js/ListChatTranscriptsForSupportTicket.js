const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all chat transcripts for a support ticket
 *
 * @summary Lists all chat transcripts for a support ticket
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListChatTranscriptsForSupportTicket.json
 */
async function listCommunicationsForANoSubscriptionSupportTicket() {
  const supportTicketName = "testticket";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const resArray = new Array();
  for await (let item of client.supportTicketChatTranscriptsNoSubscription.list(
    supportTicketName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
