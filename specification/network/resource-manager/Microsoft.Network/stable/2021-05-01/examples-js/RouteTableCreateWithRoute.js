const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or updates a route table in a specified resource group.
 *
 * @summary Create or updates a route table in a specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RouteTableCreateWithRoute.json
 */
async function createRouteTableWithRoute() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeTableName = "testrt";
  const parameters = {
    disableBgpRoutePropagation: true,
    location: "westus",
    routes: [
      {
        name: "route1",
        addressPrefix: "10.0.3.0/24",
        nextHopType: "VirtualNetworkGateway",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeTables.beginCreateOrUpdateAndWait(
    resourceGroupName,
    routeTableName,
    parameters
  );
  console.log(result);
}

createRouteTableWithRoute().catch(console.error);
