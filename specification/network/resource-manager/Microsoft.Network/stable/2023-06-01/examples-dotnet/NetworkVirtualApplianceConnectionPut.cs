using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/NetworkVirtualApplianceConnectionPut.json
// this example is just showing the usage of "NetworkVirtualApplianceConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkVirtualApplianceResource created on azure
// for more information of creating NetworkVirtualApplianceResource, please refer to the document of NetworkVirtualApplianceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkVirtualApplianceName = "nva1";
ResourceIdentifier networkVirtualApplianceResourceId = NetworkVirtualApplianceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkVirtualApplianceName);
NetworkVirtualApplianceResource networkVirtualAppliance = client.GetNetworkVirtualApplianceResource(networkVirtualApplianceResourceId);

// get the collection of this NetworkVirtualApplianceConnectionResource
NetworkVirtualApplianceConnectionCollection collection = networkVirtualAppliance.GetNetworkVirtualApplianceConnections();

// invoke the operation
string connectionName = "connection1";
NetworkVirtualApplianceConnectionData data = new NetworkVirtualApplianceConnectionData()
{
    NamePropertiesName = "connection1",
    Asn = 64512,
    TunnelIdentifier = 0,
    BgpPeerAddress =
    {
    "169.254.16.13","169.254.16.14"
    },
    EnableInternetSecurity = false,
    ConnectionRoutingConfiguration = new RoutingConfiguration()
    {
        AssociatedRouteTableId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
        PropagatedRouteTables = new PropagatedRouteTable()
        {
            Labels =
            {
            "label1"
            },
            Ids =
            {
            new WritableSubResource()
            {
            Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
            }
            },
        },
        InboundRouteMapId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
        OutboundRouteMapId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
    },
    Name = "connection1",
};
ArmOperation<NetworkVirtualApplianceConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, connectionName, data);
NetworkVirtualApplianceConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkVirtualApplianceConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
