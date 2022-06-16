const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function createATicketToRequestQuotaIncreaseForSpecificVMFamilyCoresForMachineLearningService() {
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
      "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/machine_learning_service_problemClassification_guid",
    quotaTicketDetails: {
      quotaChangeRequestSubType: "BatchAml",
      quotaChangeRequestVersion: "1.0",
      quotaChangeRequests: [
        {
          payload: '{"VMFamily":"standardA0_A7Family","NewLimit":200,"Type":"Dedicated"}',
          region: "EastUS",
        },
      ],
    },
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

createATicketToRequestQuotaIncreaseForSpecificVMFamilyCoresForMachineLearningService().catch(
  console.error
);
