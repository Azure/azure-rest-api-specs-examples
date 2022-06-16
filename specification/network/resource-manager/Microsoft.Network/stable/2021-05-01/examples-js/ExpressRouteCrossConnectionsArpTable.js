const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
 *
 * @summary Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCrossConnectionsArpTable.json
 */
async function getExpressRouteCrossConnectionsArpTable() {
  const subscriptionId = "subid";
  const resourceGroupName = "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const peeringName = "AzurePrivatePeering";
  const devicePath = "primary";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnections.beginListArpTableAndWait(
    resourceGroupName,
    crossConnectionName,
    peeringName,
    devicePath
  );
  console.log(result);
}

getExpressRouteCrossConnectionsArpTable().catch(console.error);
