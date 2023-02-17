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

// this example assumes you already have this SupportTicketCommunicationResource created on azure
// for more information of creating SupportTicketCommunicationResource, please refer to the document of SupportTicketCommunicationResource
string subscriptionId = "subid";
string supportTicketName = "testticket";
string communicationName = "testmessage";
ResourceIdentifier supportTicketCommunicationResourceId = SupportTicketCommunicationResource.CreateResourceIdentifier(subscriptionId, supportTicketName, communicationName);
SupportTicketCommunicationResource supportTicketCommunication = client.GetSupportTicketCommunicationResource(supportTicketCommunicationResourceId);

// invoke the operation
SupportTicketCommunicationResource result = await supportTicketCommunication.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportTicketCommunicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
