using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureOperatorNexus/VirtualNetworkFunctionDefinitionVersionDelete.json
// this example is just showing the usage of "NetworkFunctionDefinitionVersions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFunctionDefinitionVersionResource created on azure
// for more information of creating NetworkFunctionDefinitionVersionResource, please refer to the document of NetworkFunctionDefinitionVersionResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string publisherName = "TestPublisher";
string networkFunctionDefinitionGroupName = "TestNetworkFunctionDefinitionGroupName";
string networkFunctionDefinitionVersionName = "1.0.0";
ResourceIdentifier networkFunctionDefinitionVersionResourceId = NetworkFunctionDefinitionVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName);
NetworkFunctionDefinitionVersionResource networkFunctionDefinitionVersion = client.GetNetworkFunctionDefinitionVersionResource(networkFunctionDefinitionVersionResourceId);

// invoke the operation
await networkFunctionDefinitionVersion.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
