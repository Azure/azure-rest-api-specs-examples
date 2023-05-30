using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Communication;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/senderUsernames/get.json
// this example is just showing the usage of "SenderUsernames_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CommunicationDomainResource created on azure
// for more information of creating CommunicationDomainResource, please refer to the document of CommunicationDomainResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "contosoResourceGroup";
string emailServiceName = "contosoEmailService";
string domainName = "contoso.com";
ResourceIdentifier communicationDomainResourceId = CommunicationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, emailServiceName, domainName);
CommunicationDomainResource communicationDomainResource = client.GetCommunicationDomainResource(communicationDomainResourceId);

// get the collection of this SenderUsernameResource
SenderUsernameResourceCollection collection = communicationDomainResource.GetSenderUsernameResources();

// invoke the operation
string senderUsername = "contosoNewsAlerts";
bool result = await collection.ExistsAsync(senderUsername);

Console.WriteLine($"Succeeded: {result}");
