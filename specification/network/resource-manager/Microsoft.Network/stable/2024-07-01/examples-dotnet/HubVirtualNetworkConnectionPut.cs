using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/HubVirtualNetworkConnectionPut.json
// this example is just showing the usage of "HubVirtualNetworkConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubVirtualNetworkConnectionResource created on azure
// for more information of creating HubVirtualNetworkConnectionResource, please refer to the document of HubVirtualNetworkConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualHubName = "virtualHub1";
string connectionName = "connection1";
ResourceIdentifier hubVirtualNetworkConnectionResourceId = HubVirtualNetworkConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualHubName, connectionName);
HubVirtualNetworkConnectionResource hubVirtualNetworkConnection = client.GetHubVirtualNetworkConnectionResource(hubVirtualNetworkConnectionResourceId);

// invoke the operation
HubVirtualNetworkConnectionData data = new HubVirtualNetworkConnectionData
{
    RemoteVirtualNetworkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/SpokeVnet1"),
    EnableInternetSecurity = false,
    RoutingConfiguration = new RoutingConfiguration
    {
        AssociatedRouteTableId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
        PropagatedRouteTables = new PropagatedRouteTable
        {
            Labels = { "label1", "label2" },
            Ids = {new WritableSubResource
            {
            Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
            }},
        },
        VnetRoutes = new VnetRoute
        {
            StaticRoutesConfig = new StaticRoutesConfig
            {
                VnetLocalRouteOverrideCriteria = VnetLocalRouteOverrideCriterion.Equal,
            },
            StaticRoutes = {new StaticRoute
            {
            Name = "route1",
            AddressPrefixes = {"10.1.0.0/16", "10.2.0.0/16"},
            NextHopIPAddress = "10.0.0.68",
            }, new StaticRoute
            {
            Name = "route2",
            AddressPrefixes = {"10.3.0.0/16", "10.4.0.0/16"},
            NextHopIPAddress = "10.0.0.65",
            }},
        },
        InboundRouteMapId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
        OutboundRouteMapId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
    },
};
ArmOperation<HubVirtualNetworkConnectionResource> lro = await hubVirtualNetworkConnection.UpdateAsync(WaitUntil.Completed, data);
HubVirtualNetworkConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HubVirtualNetworkConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
