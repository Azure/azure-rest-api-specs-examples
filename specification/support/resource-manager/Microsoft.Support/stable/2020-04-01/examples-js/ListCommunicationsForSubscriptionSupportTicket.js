const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function listCommunicationsForASubscriptionSupportTicket() {
  const subscriptionId = "subid";
  const supportTicketName = "testticket";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communications.list(supportTicketName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCommunicationsForASubscriptionSupportTicket().catch(console.error);
