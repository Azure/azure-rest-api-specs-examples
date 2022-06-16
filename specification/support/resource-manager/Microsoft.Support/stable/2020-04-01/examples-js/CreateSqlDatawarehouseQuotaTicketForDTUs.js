const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function createATicketToRequestQuotaIncreaseForDtUsForAzureSynapseAnalytics() {
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
      "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/sql_datawarehouse_problemClassification_guid",
    quotaTicketDetails: {
      quotaChangeRequestSubType: "DTUs",
      quotaChangeRequestVersion: "1.0",
      quotaChangeRequests: [
        {
          payload: '{"ServerName":"testserver","NewLimit":54000}',
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

createATicketToRequestQuotaIncreaseForDtUsForAzureSynapseAnalytics().catch(console.error);
