const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This API allows you to update the severity level, ticket status, and your contact information in the support ticket.<br/><br/>Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an Azure support engineer. In such a case, contact your support engineer to request severity update by adding a new communication using the Communications API.
 *
 * @summary This API allows you to update the severity level, ticket status, and your contact information in the support ticket.<br/><br/>Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an Azure support engineer. In such a case, contact your support engineer to request severity update by adding a new communication using the Communications API.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/UpdateSeverityOfSupportTicket.json
 */
async function updateSeverityOfASupportTicket() {
  const supportTicketName = "testticket";
  const updateSupportTicket = { severity: "critical" };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.supportTicketsNoSubscription.update(
    supportTicketName,
    updateSupportTicket
  );
  console.log(result);
}
