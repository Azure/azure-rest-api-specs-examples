using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/L3Networks_Create.json
// this example is just showing the usage of "L3Networks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this L3NetworkResource
L3NetworkCollection collection = resourceGroupResource.GetL3Networks();

// invoke the operation
string l3NetworkName = "l3NetworkName";
L3NetworkData data = new L3NetworkData(new AzureLocation("location"), new ExtendedLocation("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName", "CustomLocation"), "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/l3IsolationDomainName", 12)
{
    InterfaceName = "eth0",
    IPAllocationType = IPAllocationType.DualStack,
    IPv4ConnectedPrefix = "198.51.100.0/24",
    IPv6ConnectedPrefix = "2001:db8::/64",
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2",
    },
};
ArmOperation<L3NetworkResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, l3NetworkName, data);
L3NetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
L3NetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
