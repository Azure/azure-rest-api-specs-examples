using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Communication;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/smtpUsername/createOrUpdate.json
// this example is just showing the usage of "SmtpUsernames_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CommunicationSmtpUsernameResource created on azure
// for more information of creating CommunicationSmtpUsernameResource, please refer to the document of CommunicationSmtpUsernameResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "contosoResourceGroup";
string communicationServiceName = "contosoACSService";
string smtpUsername = "smtpusername1";
ResourceIdentifier communicationSmtpUsernameResourceId = CommunicationSmtpUsernameResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, communicationServiceName, smtpUsername);
CommunicationSmtpUsernameResource communicationSmtpUsername = client.GetCommunicationSmtpUsernameResource(communicationSmtpUsernameResourceId);

// invoke the operation
CommunicationSmtpUsernameData data = new CommunicationSmtpUsernameData
{
    Username = "newuser1@contoso.com",
    EntraApplicationId = "aaaa1111-bbbb-2222-3333-aaaa111122bb",
    TenantId = Guid.Parse("aaaa1111-bbbb-2222-3333-aaaa11112222"),
};
ArmOperation<CommunicationSmtpUsernameResource> lro = await communicationSmtpUsername.UpdateAsync(WaitUntil.Completed, data);
CommunicationSmtpUsernameResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CommunicationSmtpUsernameData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
