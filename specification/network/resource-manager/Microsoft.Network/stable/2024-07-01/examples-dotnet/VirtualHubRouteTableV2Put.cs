using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualHubRouteTableV2Put.json
// this example is just showing the usage of "VirtualHubRouteTableV2s_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualHubRouteTableV2Resource created on azure
// for more information of creating VirtualHubRouteTableV2Resource, please refer to the document of VirtualHubRouteTableV2Resource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "virtualHub1";
string routeTableName = "virtualHubRouteTable1a";
ResourceIdentifier virtualHubRouteTableV2ResourceId = VirtualHubRouteTableV2Resource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, routeTableName);
VirtualHubRouteTableV2Resource virtualHubRouteTableV2 = client.GetVirtualHubRouteTableV2Resource(virtualHubRouteTableV2ResourceId);

// invoke the operation
VirtualHubRouteTableV2Data data = new VirtualHubRouteTableV2Data
{
    Routes = {new VirtualHubRouteV2
    {
    DestinationType = "CIDR",
    Destinations = {"20.10.0.0/16", "20.20.0.0/16"},
    NextHopType = "IPAddress",
    NextHops = {"10.0.0.68"},
    }, new VirtualHubRouteV2
    {
    DestinationType = "CIDR",
    Destinations = {"0.0.0.0/0"},
    NextHopType = "IPAddress",
    NextHops = {"10.0.0.68"},
    }},
    AttachedConnections = { "All_Vnets" },
};
ArmOperation<VirtualHubRouteTableV2Resource> lro = await virtualHubRouteTableV2.UpdateAsync(WaitUntil.Completed, data);
VirtualHubRouteTableV2Resource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualHubRouteTableV2Data resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
