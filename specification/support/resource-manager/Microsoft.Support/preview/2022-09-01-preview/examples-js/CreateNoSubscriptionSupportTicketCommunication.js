const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds a new customer communication to an Azure support ticket.
 *
 * @summary Adds a new customer communication to an Azure support ticket.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateNoSubscriptionSupportTicketCommunication.json
 */
async function addCommunicationToNoSubscriptionTicket() {
  const supportTicketName = "testticket";
  const communicationName = "testcommunication";
  const createCommunicationParameters = {
    body: "This is a test message from a customer!",
    sender: "user@contoso.com",
    subject: "This is a test message from a customer!",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.communicationsNoSubscription.beginCreateAndWait(
    supportTicketName,
    communicationName,
    createCommunicationParameters
  );
  console.log(result);
}
