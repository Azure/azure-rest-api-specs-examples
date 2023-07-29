using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpCommunities_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "IpCommunities_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricIPCommunityResource created on azure
// for more information of creating NetworkFabricIPCommunityResource, please refer to the document of NetworkFabricIPCommunityResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string ipCommunityName = "example-ipcommunity";
ResourceIdentifier networkFabricIPCommunityResourceId = NetworkFabricIPCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ipCommunityName);
NetworkFabricIPCommunityResource networkFabricIPCommunity = client.GetNetworkFabricIPCommunityResource(networkFabricIPCommunityResourceId);

// invoke the operation
await networkFabricIPCommunity.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
