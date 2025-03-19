using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.VoiceServices.Models;
using Azure.ResourceManager.VoiceServices;

// Generated from example definition: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/CommunicationsGateways_ListByResourceGroup.json
// this example is just showing the usage of "CommunicationsGateways_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this VoiceServicesCommunicationsGatewayResource
VoiceServicesCommunicationsGatewayCollection collection = resourceGroupResource.GetVoiceServicesCommunicationsGateways();

// invoke the operation and iterate over the result
await foreach (VoiceServicesCommunicationsGatewayResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    VoiceServicesCommunicationsGatewayData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
