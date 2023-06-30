using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricControllers_enableWorkloadManagementNetwork_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkFabricControllers_enableWorkloadManagementNetwork" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricControllerResource created on azure
// for more information of creating NetworkFabricControllerResource, please refer to the document of NetworkFabricControllerResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string networkFabricControllerName = "networkFabricControllerName";
ResourceIdentifier networkFabricControllerResourceId = NetworkFabricControllerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFabricControllerName);
NetworkFabricControllerResource networkFabricController = client.GetNetworkFabricControllerResource(networkFabricControllerResourceId);

// invoke the operation
await networkFabricController.EnableWorkloadManagementNetworkAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
