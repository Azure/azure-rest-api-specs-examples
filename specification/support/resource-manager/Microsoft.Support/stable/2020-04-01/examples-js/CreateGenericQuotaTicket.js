const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function createATicketToRequestQuotaIncreaseForServicesThatDoNotRequireAdditionalDetailsInTheQuotaTicketDetailsObject() {
  const subscriptionId = "subid";
  const supportTicketName = "testticket";
  const createSupportTicketParameters = {
    description: "Increase the maximum throughput per container limit to 10000 for account foo bar",
    contactDetails: {
      country: "usa",
      firstName: "abc",
      lastName: "xyz",
      preferredContactMethod: "email",
      preferredSupportLanguage: "en-US",
      preferredTimeZone: "Pacific Standard Time",
      primaryEmailAddress: "abc@contoso.com",
    },
    problemClassificationId:
      "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cosmosdb_problemClassification_guid",
    serviceId: "/providers/Microsoft.Support/services/quota_service_guid",
    severity: "moderate",
    title: "my title",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.supportTickets.beginCreateAndWait(
    supportTicketName,
    createSupportTicketParameters
  );
  console.log(result);
}

createATicketToRequestQuotaIncreaseForServicesThatDoNotRequireAdditionalDetailsInTheQuotaTicketDetailsObject().catch(
  console.error
);
