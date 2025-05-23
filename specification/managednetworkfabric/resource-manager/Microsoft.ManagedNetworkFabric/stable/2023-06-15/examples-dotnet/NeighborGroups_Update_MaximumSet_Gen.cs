using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NeighborGroups_Update_MaximumSet_Gen.json
// this example is just showing the usage of "NeighborGroups_Update" operation, for the dependent resources, they will have to be created separately.

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
NetworkFabricNeighborGroupPatch patch = new NetworkFabricNeighborGroupPatch
{
    Annotation = "Updating",
    Destination = new NeighborGroupDestination
    {
        IPv4Addresses = { IPAddress.Parse("10.10.10.10"), IPAddress.Parse("20.10.10.10"), IPAddress.Parse("30.10.10.10"), IPAddress.Parse("40.10.10.10"), IPAddress.Parse("50.10.10.10"), IPAddress.Parse("60.10.10.10"), IPAddress.Parse("70.10.10.10"), IPAddress.Parse("80.10.10.10"), IPAddress.Parse("90.10.10.10") },
        IPv6Addresses = { "2F::/100" },
    },
    Tags =
    {
    ["key8107"] = "2345"
    },
};
ArmOperation<NetworkFabricNeighborGroupResource> lro = await networkFabricNeighborGroup.UpdateAsync(WaitUntil.Completed, patch);
NetworkFabricNeighborGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricNeighborGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
