using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2021-02-01/examples/RouteTableCreateWithRoute.json
// this example is just showing the usage of "RouteTables_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this RouteTableResource
RouteTableCollection collection = resourceGroupResource.GetRouteTables();

// invoke the operation
string routeTableName = "testrt";
RouteTableData data = new RouteTableData()
{
    Routes =
    {
    new RouteData()
    {
    AddressPrefix = "10.0.3.0/24",
    NextHopType = RouteNextHopType.VirtualNetworkGateway,
    Name = "route1",
    }
    },
    DisableBgpRoutePropagation = true,
    Location = new AzureLocation("westus"),
};
ArmOperation<RouteTableResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, routeTableName, data);
RouteTableResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RouteTableData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
