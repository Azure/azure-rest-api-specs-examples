using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateSupportTicketCommunication.json
// this example is just showing the usage of "Communications_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SupportTicketCommunicationResource created on azure
// for more information of creating SupportTicketCommunicationResource, please refer to the document of SupportTicketCommunicationResource
string subscriptionId = "132d901f-189d-4381-9214-fe68e27e05a1";
string supportTicketName = "testticket";
string communicationName = "testcommunication";
ResourceIdentifier supportTicketCommunicationResourceId = SupportTicketCommunicationResource.CreateResourceIdentifier(subscriptionId, supportTicketName, communicationName);
SupportTicketCommunicationResource supportTicketCommunication = client.GetSupportTicketCommunicationResource(supportTicketCommunicationResourceId);

// invoke the operation
SupportTicketCommunicationData data = new SupportTicketCommunicationData("This is a test message from a customer!", "This is a test message from a customer!")
{
    Sender = "user@contoso.com",
};
ArmOperation<SupportTicketCommunicationResource> lro = await supportTicketCommunication.UpdateAsync(WaitUntil.Completed, data);
SupportTicketCommunicationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportTicketCommunicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
