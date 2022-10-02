const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the effective routes configured for the Virtual Hub resource or the specified resource .
 *
 * @summary Gets the effective routes configured for the Virtual Hub resource or the specified resource .
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/EffectiveRoutesListForRouteTable.json
 */
async function effectiveRoutesForARouteTableResource() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const effectiveRoutesParameters = {
    resourceId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1",
    virtualWanResourceType: "RouteTable",
  };
  const options = {
    effectiveRoutesParameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubs.beginGetEffectiveVirtualHubRoutesAndWait(
    resourceGroupName,
    virtualHubName,
    options
  );
  console.log(result);
}

effectiveRoutesForARouteTableResource().catch(console.error);
