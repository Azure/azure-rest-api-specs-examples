const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new support ticket for Billing, and Subscription Management issues. Learn the [prerequisites](https://aka.ms/supportAPI) required to create a support ticket.<br/><br/>Always call the Services and ProblemClassifications API to get the most recent set of services and problem categories required for support ticket creation.<br/><br/>Adding attachments is not currently supported via the API. To add a file to an existing support ticket, visit the [Manage support ticket](https://portal.azure.com/#blade/Microsoft_Azure_Support/HelpAndSupportBlade/managesupportrequest) page in the Azure portal, select the support ticket, and use the file upload control to add a new file.<br/><br/>Providing consent to share diagnostic information with Azure support is currently not supported via the API. The Azure support engineer working on your ticket will reach out to you for consent if your issue requires gathering diagnostic information from your Azure resources.<br/><br/>
 *
 * @summary Creates a new support ticket for Billing, and Subscription Management issues. Learn the [prerequisites](https://aka.ms/supportAPI) required to create a support ticket.<br/><br/>Always call the Services and ProblemClassifications API to get the most recent set of services and problem categories required for support ticket creation.<br/><br/>Adding attachments is not currently supported via the API. To add a file to an existing support ticket, visit the [Manage support ticket](https://portal.azure.com/#blade/Microsoft_Azure_Support/HelpAndSupportBlade/managesupportrequest) page in the Azure portal, select the support ticket, and use the file upload control to add a new file.<br/><br/>Providing consent to share diagnostic information with Azure support is currently not supported via the API. The Azure support engineer working on your ticket will reach out to you for consent if your issue requires gathering diagnostic information from your Azure resources.<br/><br/>
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateBillingSupportTicket.json
 */
async function createATicketForBillingRelatedIssues() {
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
    fileWorkspaceName: "6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066",
    problemClassificationId:
      "/providers/Microsoft.Support/services/billing_service_guid/problemClassifications/billing_problemClassification_guid",
    serviceId: "/providers/Microsoft.Support/services/billing_service_guid",
    severity: "moderate",
    supportPlanId:
      "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=",
    title: "my title",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const result = await client.supportTicketsNoSubscription.beginCreateAndWait(
    supportTicketName,
    createSupportTicketParameters
  );
  console.log(result);
}
