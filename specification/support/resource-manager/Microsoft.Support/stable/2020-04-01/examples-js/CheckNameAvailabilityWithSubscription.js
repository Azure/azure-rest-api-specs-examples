const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function checksWhetherNameIsAvailableForSupportTicketResource() {
  const subscriptionId = "subid";
  const checkNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Support/supportTickets",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.supportTickets.checkNameAvailability(checkNameAvailabilityInput);
  console.log(result);
}

checksWhetherNameIsAvailableForSupportTicketResource().catch(console.error);
