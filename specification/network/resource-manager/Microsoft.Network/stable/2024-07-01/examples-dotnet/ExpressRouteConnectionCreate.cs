using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ExpressRouteConnectionCreate.json
// this example is just showing the usage of "ExpressRouteConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteGatewayResource created on azure
// for more information of creating ExpressRouteGatewayResource, please refer to the document of ExpressRouteGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "resourceGroupName";
string expressRouteGatewayName = "gateway-2";
ResourceIdentifier expressRouteGatewayResourceId = ExpressRouteGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, expressRouteGatewayName);
ExpressRouteGatewayResource expressRouteGateway = client.GetExpressRouteGatewayResource(expressRouteGatewayResourceId);

// get the collection of this ExpressRouteConnectionResource
ExpressRouteConnectionCollection collection = expressRouteGateway.GetExpressRouteConnections();

// invoke the operation
string connectionName = "connectionName";
ExpressRouteConnectionData data = new ExpressRouteConnectionData
{
    ExpressRouteCircuitPeeringId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
    AuthorizationKey = "authorizationKey",
    RoutingWeight = 2,
    RoutingConfiguration = new RoutingConfiguration
    {
        AssociatedRouteTableId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
        PropagatedRouteTables = new PropagatedRouteTable
        {
            Labels = { "label1", "label2" },
            Ids = {new WritableSubResource
            {
            Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
            }, new WritableSubResource
            {
            Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
            }, new WritableSubResource
            {
            Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"),
            }},
        },
        InboundRouteMapId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
        OutboundRouteMapId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
    },
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName"),
    Name = "connectionName",
};
ArmOperation<ExpressRouteConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, connectionName, data);
ExpressRouteConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExpressRouteConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
