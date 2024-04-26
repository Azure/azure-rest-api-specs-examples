const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all chat transcripts for a support ticket
 *
 * @summary Lists all chat transcripts for a support ticket
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListChatTranscriptsForSupportTicket.json
 */
async function listChatTranscriptsForANoSubscriptionSupportTicket() {
  const supportTicketName = "testticket";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const resArray = new Array();
  for await (let item of client.chatTranscriptsNoSubscription.list(supportTicketName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
