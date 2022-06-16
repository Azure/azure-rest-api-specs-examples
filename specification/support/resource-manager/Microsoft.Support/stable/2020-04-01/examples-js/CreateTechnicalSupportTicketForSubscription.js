const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function createATicketForTechnicalIssueRelatedToASpecificResource() {
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
      "/providers/Microsoft.Support/services/virtual_machine_running_linux_service_guid/problemClassifications/problemClassification_guid",
    serviceId: "/providers/Microsoft.Support/services/cddd3eb5-1830-b494-44fd-782f691479dc",
    severity: "moderate",
    technicalTicketDetails: {
      resourceId:
        "/subscriptions/subid/resourceGroups/test/providers/Microsoft.Compute/virtualMachines/testserver",
    },
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

createATicketForTechnicalIssueRelatedToASpecificResource().catch(console.error);
