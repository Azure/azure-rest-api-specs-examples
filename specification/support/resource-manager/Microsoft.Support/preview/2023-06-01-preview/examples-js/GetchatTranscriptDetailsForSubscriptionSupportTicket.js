const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns chatTranscript details for a support ticket under a subscription.
 *
 * @summary Returns chatTranscript details for a support ticket under a subscription.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/GetchatTranscriptDetailsForSubscriptionSupportTicket.json
 */
async function getChatTranscriptDetailsForASubscriptionSupportTicket() {
  const subscriptionId = process.env["SUPPORT_SUBSCRIPTION_ID"] || "subid";
  const supportTicketName = "testticket";
  const chatTranscriptName = "69586795-45e9-45b5-bd9e-c9bb237d3e44";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.chatTranscripts.get(supportTicketName, chatTranscriptName);
  console.log(result);
}
