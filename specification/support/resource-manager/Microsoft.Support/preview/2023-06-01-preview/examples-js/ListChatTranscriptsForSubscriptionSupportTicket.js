const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all chat transcripts for a support ticket under subscription
 *
 * @summary Lists all chat transcripts for a support ticket under subscription
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ListChatTranscriptsForSubscriptionSupportTicket.json
 */
async function listChatTranscriptsForASubscriptionSupportTicket() {
  const subscriptionId = process.env["SUPPORT_SUBSCRIPTION_ID"] || "subid";
  const supportTicketName = "testticket";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.chatTranscripts.list(supportTicketName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
