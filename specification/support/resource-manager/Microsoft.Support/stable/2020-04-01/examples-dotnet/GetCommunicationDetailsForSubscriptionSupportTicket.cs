using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/GetCommunicationDetailsForSubscriptionSupportTicket.json
// this example is just showing the usage of "Communications_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SupportTicketResource created on azure
// for more information of creating SupportTicketResource, please refer to the document of SupportTicketResource
string subscriptionId = "subid";
string supportTicketName = "testticket";
ResourceIdentifier supportTicketResourceId = SupportTicketResource.CreateResourceIdentifier(subscriptionId, supportTicketName);
SupportTicketResource supportTicket = client.GetSupportTicketResource(supportTicketResourceId);

// get the collection of this SupportTicketCommunicationResource
SupportTicketCommunicationCollection collection = supportTicket.GetSupportTicketCommunications();

// invoke the operation
string communicationName = "testmessage";
bool result = await collection.ExistsAsync(communicationName);

Console.WriteLine($"Succeeded: {result}");
