using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Support;
using Azure.ResourceManager.Support.Models;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CreateGenericQuotaTicket.json
// this example is just showing the usage of "SupportTickets_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subid";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this SupportTicketResource
SupportTicketCollection collection = subscriptionResource.GetSupportTickets();

// invoke the operation
string supportTicketName = "testticket";
SupportTicketData data = new SupportTicketData()
{
    Description = "Increase the maximum throughput per container limit to 10000 for account foo bar",
    ProblemClassificationId = "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cosmosdb_problemClassification_guid",
    Severity = SupportSeverityLevel.Moderate,
    ContactDetails = new SupportContactProfile("abc", "xyz", PreferredContactMethod.Email, "abc@contoso.com", "Pacific Standard Time", "usa", "en-US"),
    Title = "my title",
    ServiceId = "/providers/Microsoft.Support/services/quota_service_guid",
};
ArmOperation<SupportTicketResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, supportTicketName, data);
SupportTicketResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportTicketData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
