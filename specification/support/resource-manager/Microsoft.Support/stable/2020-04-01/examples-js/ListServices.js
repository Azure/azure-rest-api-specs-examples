const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function getsListOfServicesForWhichASupportTicketCanBeCreated() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.services.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsListOfServicesForWhichASupportTicketCanBeCreated().catch(console.error);
