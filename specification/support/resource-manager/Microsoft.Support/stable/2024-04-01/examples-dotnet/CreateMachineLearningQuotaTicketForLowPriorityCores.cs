using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Support.Models;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateMachineLearningQuotaTicketForLowPriorityCores.json
// this example is just showing the usage of "SupportTickets_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "132d901f-189d-4381-9214-fe68e27e05a1";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this SubscriptionSupportTicketResource
SubscriptionSupportTicketCollection collection = subscriptionResource.GetSubscriptionSupportTickets();

// invoke the operation
string supportTicketName = "testticket";
SupportTicketData data = new SupportTicketData(
    "my description",
    "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/machine_learning_service_problemClassification_guid",
    SupportSeverityLevel.Moderate,
    AdvancedDiagnosticConsent.Yes,
    new SupportContactProfile(
    "abc",
    "xyz",
    PreferredContactMethod.Email,
    "abc@contoso.com",
    "Pacific Standard Time",
    "usa",
    "en-US"),
    "my title",
    "/providers/Microsoft.Support/services/quota_service_guid")
{
    SupportPlanId = "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=",
    QuotaTicketDetails = new QuotaTicketDetails
    {
        QuotaChangeRequestSubType = "BatchAml",
        QuotaChangeRequestVersion = "1.0",
        QuotaChangeRequests = {new SupportQuotaChangeContent
        {
        Region = "EastUS",
        Payload = "{\"NewLimit\":200,\"Type\":\"LowPriority\"}",
        }},
    },
};
ArmOperation<SubscriptionSupportTicketResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, supportTicketName, data);
SubscriptionSupportTicketResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportTicketData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
