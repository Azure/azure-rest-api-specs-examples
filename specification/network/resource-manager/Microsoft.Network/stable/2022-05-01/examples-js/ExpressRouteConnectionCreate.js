const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
 *
 * @summary Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteConnectionCreate.json
 */
async function expressRouteConnectionCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "resourceGroupName";
  const expressRouteGatewayName = "gateway-2";
  const connectionName = "connectionName";
  const putExpressRouteConnectionParameters = {
    name: "connectionName",
    authorizationKey: "authorizationKey",
    expressRouteCircuitPeering: {
      id: "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering",
    },
    id: "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName",
    routingConfiguration: {
      associatedRouteTable: {
        id: "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1",
      },
      inboundRouteMap: {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1",
      },
      outboundRouteMap: {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2",
      },
      propagatedRouteTables: {
        ids: [
          {
            id: "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1",
          },
          {
            id: "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2",
          },
          {
            id: "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3",
          },
        ],
        labels: ["label1", "label2"],
      },
    },
    routingWeight: 2,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    expressRouteGatewayName,
    connectionName,
    putExpressRouteConnectionParameters
  );
  console.log(result);
}

expressRouteConnectionCreate().catch(console.error);
