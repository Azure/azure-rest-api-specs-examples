using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/CreateNoSubscriptionSupportTicketCommunication.json
// this example is just showing the usage of "CommunicationsNoSubscription_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantSupportTicketResource created on azure
// for more information of creating TenantSupportTicketResource, please refer to the document of TenantSupportTicketResource
string supportTicketName = "testticket";
ResourceIdentifier tenantSupportTicketResourceId = TenantSupportTicketResource.CreateResourceIdentifier(supportTicketName);
TenantSupportTicketResource tenantSupportTicket = client.GetTenantSupportTicketResource(tenantSupportTicketResourceId);

// get the collection of this SupportTicketNoSubCommunicationResource
SupportTicketNoSubCommunicationCollection collection = tenantSupportTicket.GetSupportTicketNoSubCommunications();

// invoke the operation
string communicationName = "testcommunication";
SupportTicketCommunicationData data = new SupportTicketCommunicationData()
{
    Sender = "user@contoso.com",
    Subject = "This is a test message from a customer!",
    Body = "This is a test message from a customer!",
};
ArmOperation<SupportTicketNoSubCommunicationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, communicationName, data);
SupportTicketNoSubCommunicationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportTicketCommunicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
