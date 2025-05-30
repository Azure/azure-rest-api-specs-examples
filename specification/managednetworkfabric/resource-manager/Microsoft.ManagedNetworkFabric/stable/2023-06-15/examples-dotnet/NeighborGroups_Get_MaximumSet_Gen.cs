using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NeighborGroups_Get_MaximumSet_Gen.json
// this example is just showing the usage of "NeighborGroups_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricNeighborGroupResource created on azure
// for more information of creating NetworkFabricNeighborGroupResource, please refer to the document of NetworkFabricNeighborGroupResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string neighborGroupName = "example-neighborGroup";
ResourceIdentifier networkFabricNeighborGroupResourceId = NetworkFabricNeighborGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, neighborGroupName);
NetworkFabricNeighborGroupResource networkFabricNeighborGroup = client.GetNetworkFabricNeighborGroupResource(networkFabricNeighborGroupResourceId);

// invoke the operation
NetworkFabricNeighborGroupResource result = await networkFabricNeighborGroup.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricNeighborGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
