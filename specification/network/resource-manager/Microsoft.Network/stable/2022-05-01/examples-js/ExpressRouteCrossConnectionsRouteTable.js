const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the currently advertised routes table associated with the express route cross connection in a resource group.
 *
 * @summary Gets the currently advertised routes table associated with the express route cross connection in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteCrossConnectionsRouteTable.json
 */
async function getExpressRouteCrossConnectionsRouteTable() {
  const subscriptionId = "subid";
  const resourceGroupName = "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const peeringName = "AzurePrivatePeering";
  const devicePath = "primary";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnections.beginListRoutesTableAndWait(
    resourceGroupName,
    crossConnectionName,
    peeringName,
    devicePath
  );
  console.log(result);
}

getExpressRouteCrossConnectionsRouteTable().catch(console.error);
