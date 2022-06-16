const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function createATicketToRequestQuotaIncreaseForLowPriorityCoresForABatchAccount() {
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
      "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/batch_problemClassification_guid",
    quotaTicketDetails: {
      quotaChangeRequestSubType: "Account",
      quotaChangeRequestVersion: "1.0",
      quotaChangeRequests: [
        {
          payload: '{"AccountName":"test","NewLimit":200,"Type":"LowPriority"}',
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

createATicketToRequestQuotaIncreaseForLowPriorityCoresForABatchAccount().catch(console.error);
