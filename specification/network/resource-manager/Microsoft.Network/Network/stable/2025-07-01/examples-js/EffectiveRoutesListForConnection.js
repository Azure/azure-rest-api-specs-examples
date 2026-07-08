const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the effective routes configured for the Virtual Hub resource or the specified resource .
 *
 * @summary gets the effective routes configured for the Virtual Hub resource or the specified resource .
 * x-ms-original-file: 2025-07-01/EffectiveRoutesListForConnection.json
 */
async function effectiveRoutesForAConnectionResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubs.getEffectiveVirtualHubRoutes("rg1", "virtualHub1", {
    effectiveRoutesParameters: {
      resourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName",
      virtualWanResourceType: "ExpressRouteConnection",
    },
  });
  console.log(result);
}
