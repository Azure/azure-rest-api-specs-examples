const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function createATicketForBillingRelatedIssues() {
  const subscriptionId = "subid";
  const supportTicketName = "testticket";
  const createSupportTicketParameters = {
    description: "my description",
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
      "/providers/Microsoft.Support/services/billing_service_guid/problemClassifications/billing_problemClassification_guid",
    serviceId: "/providers/Microsoft.Support/services/billing_service_guid",
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

createATicketForBillingRelatedIssues().catch(console.error);
