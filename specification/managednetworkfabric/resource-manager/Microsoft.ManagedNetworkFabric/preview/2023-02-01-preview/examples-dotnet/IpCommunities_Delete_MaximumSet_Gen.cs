using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpCommunities_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "IpCommunities_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IPCommunityResource created on azure
// for more information of creating IPCommunityResource, please refer to the document of IPCommunityResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string ipCommunityName = "IpCommunityList1";
ResourceIdentifier ipCommunityResourceId = IPCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ipCommunityName);
IPCommunityResource ipCommunity = client.GetIPCommunityResource(ipCommunityResourceId);

// invoke the operation
await ipCommunity.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
