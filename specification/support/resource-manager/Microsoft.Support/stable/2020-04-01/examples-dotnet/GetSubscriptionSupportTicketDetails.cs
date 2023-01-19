using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Support;
using Azure.ResourceManager.Support.Models;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/GetSubscriptionSupportTicketDetails.json
// this example is just showing the usage of "SupportTickets_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
SupportTicketResource result = await supportTicket.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportTicketData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
