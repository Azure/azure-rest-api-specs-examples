const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a RouteMap.
 *
 * @summary Deletes a RouteMap.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/RouteMapDelete.json
 */
async function routeMapDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const routeMapName = "routeMap1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeMaps.beginDeleteAndWait(
    resourceGroupName,
    virtualHubName,
    routeMapName
  );
  console.log(result);
}

routeMapDelete().catch(console.error);
