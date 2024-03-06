using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/GetchatTranscriptDetailsForSupportTicket.json
// this example is just showing the usage of "ChatTranscriptsNoSubscription_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantSupportTicketResource created on azure
// for more information of creating TenantSupportTicketResource, please refer to the document of TenantSupportTicketResource
string supportTicketName = "testticket";
ResourceIdentifier tenantSupportTicketResourceId = TenantSupportTicketResource.CreateResourceIdentifier(supportTicketName);
TenantSupportTicketResource tenantSupportTicket = client.GetTenantSupportTicketResource(tenantSupportTicketResourceId);

// get the collection of this SupportTicketNoSubChatTranscriptResource
SupportTicketNoSubChatTranscriptCollection collection = tenantSupportTicket.GetSupportTicketNoSubChatTranscripts();

// invoke the operation
string chatTranscriptName = "b371192a-b094-4a71-b093-7246029b0a54";
NullableResponse<SupportTicketNoSubChatTranscriptResource> response = await collection.GetIfExistsAsync(chatTranscriptName);
SupportTicketNoSubChatTranscriptResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ChatTranscriptDetailData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
