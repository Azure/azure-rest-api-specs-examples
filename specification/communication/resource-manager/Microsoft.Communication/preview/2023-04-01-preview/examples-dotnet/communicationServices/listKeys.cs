using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Communication;
using Azure.ResourceManager.Communication.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/communicationServices/listKeys.json
// this example is just showing the usage of "CommunicationServices_ListKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CommunicationServiceResource created on azure
// for more information of creating CommunicationServiceResource, please refer to the document of CommunicationServiceResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "MyResourceGroup";
string communicationServiceName = "MyCommunicationResource";
ResourceIdentifier communicationServiceResourceId = CommunicationServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, communicationServiceName);
CommunicationServiceResource communicationServiceResource = client.GetCommunicationServiceResource(communicationServiceResourceId);

// invoke the operation
CommunicationServiceKeys result = await communicationServiceResource.GetKeysAsync();

Console.WriteLine($"Succeeded: {result}");
