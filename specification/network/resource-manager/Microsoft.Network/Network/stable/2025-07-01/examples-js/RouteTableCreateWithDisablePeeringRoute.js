const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or updates a route table in a specified resource group.
 *
 * @summary create or updates a route table in a specified resource group.
 * x-ms-original-file: 2025-07-01/RouteTableCreateWithDisablePeeringRoute.json
 */
async function createRouteTableWithDisablePeeringRoute() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeTables.createOrUpdate("rg1", "testrt", {
    location: "westus",
    disableBgpRoutePropagation: true,
    disablePeeringRoute: "All",
  });
  console.log(result);
}
