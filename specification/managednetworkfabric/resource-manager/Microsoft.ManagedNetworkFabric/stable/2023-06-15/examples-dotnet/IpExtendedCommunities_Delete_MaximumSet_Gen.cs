using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "IpExtendedCommunities_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricIPExtendedCommunityResource created on azure
// for more information of creating NetworkFabricIPExtendedCommunityResource, please refer to the document of NetworkFabricIPExtendedCommunityResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string ipExtendedCommunityName = "example-ipExtendedCommunity";
ResourceIdentifier networkFabricIPExtendedCommunityResourceId = NetworkFabricIPExtendedCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ipExtendedCommunityName);
NetworkFabricIPExtendedCommunityResource networkFabricIPExtendedCommunity = client.GetNetworkFabricIPExtendedCommunityResource(networkFabricIPExtendedCommunityResourceId);

// invoke the operation
await networkFabricIPExtendedCommunity.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
