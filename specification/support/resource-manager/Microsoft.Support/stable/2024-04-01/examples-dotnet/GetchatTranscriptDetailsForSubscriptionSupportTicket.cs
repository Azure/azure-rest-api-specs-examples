using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/GetchatTranscriptDetailsForSubscriptionSupportTicket.json
// this example is just showing the usage of "ChatTranscripts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SupportTicketChatTranscriptResource created on azure
// for more information of creating SupportTicketChatTranscriptResource, please refer to the document of SupportTicketChatTranscriptResource
string subscriptionId = "132d901f-189d-4381-9214-fe68e27e05a1";
string supportTicketName = "testticket";
string chatTranscriptName = "69586795-45e9-45b5-bd9e-c9bb237d3e44";
ResourceIdentifier supportTicketChatTranscriptResourceId = SupportTicketChatTranscriptResource.CreateResourceIdentifier(subscriptionId, supportTicketName, chatTranscriptName);
SupportTicketChatTranscriptResource supportTicketChatTranscript = client.GetSupportTicketChatTranscriptResource(supportTicketChatTranscriptResourceId);

// invoke the operation
SupportTicketChatTranscriptResource result = await supportTicketChatTranscript.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ChatTranscriptDetailData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
