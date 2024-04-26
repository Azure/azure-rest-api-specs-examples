const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This API allows you to update the severity level, ticket status, advanced diagnostic consent and your contact information in the support ticket.<br/><br/>Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an Azure support engineer. In such a case, contact your support engineer to request severity update by adding a new communication using the Communications API.
 *
 * @summary This API allows you to update the severity level, ticket status, advanced diagnostic consent and your contact information in the support ticket.<br/><br/>Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an Azure support engineer. In such a case, contact your support engineer to request severity update by adding a new communication using the Communications API.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/UpdateAdvancedDiagnosticConsentOfSupportTicketForSubscription.json
 */
async function updateAdvancedDiagnosticConsentOfASubscriptionSupportTicket() {
  const subscriptionId =
    process.env["SUPPORT_SUBSCRIPTION_ID"] || "132d901f-189d-4381-9214-fe68e27e05a1";
  const supportTicketName = "testticket";
  const updateSupportTicket = {
    advancedDiagnosticConsent: "Yes",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.supportTickets.update(supportTicketName, updateSupportTicket);
  console.log(result);
}
