using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Support;
using Azure.ResourceManager.Support.Models;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CheckNameAvailabilityForSupportTicketCommunication.json
// this example is just showing the usage of "Communications_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

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
SupportNameAvailabilityContent content = new SupportNameAvailabilityContent("sampleName", SupportResourceType.MicrosoftSupportCommunications);
SupportNameAvailabilityResult result = await supportTicket.CheckCommunicationNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
