const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function checksWhetherNameIsAvailableForCommunicationResource() {
  const subscriptionId = "subid";
  const supportTicketName = "testticket";
  const checkNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Support/communications",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.communications.checkNameAvailability(
    supportTicketName,
    checkNameAvailabilityInput
  );
  console.log(result);
}

checksWhetherNameIsAvailableForCommunicationResource().catch(console.error);
