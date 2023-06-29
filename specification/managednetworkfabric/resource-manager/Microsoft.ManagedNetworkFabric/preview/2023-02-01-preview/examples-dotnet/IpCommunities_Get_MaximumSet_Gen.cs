using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpCommunities_Get_MaximumSet_Gen.json
// this example is just showing the usage of "IpCommunities_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IPCommunityResource created on azure
// for more information of creating IPCommunityResource, please refer to the document of IPCommunityResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string ipCommunityName = "example-ipCommunity";
ResourceIdentifier ipCommunityResourceId = IPCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ipCommunityName);
IPCommunityResource ipCommunity = client.GetIPCommunityResource(ipCommunityResourceId);

// invoke the operation
IPCommunityResource result = await ipCommunity.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IPCommunityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
