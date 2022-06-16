const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all peerings in a specified ExpressRouteCrossConnection.
 *
 * @summary Gets all peerings in a specified ExpressRouteCrossConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCrossConnectionBgpPeeringList.json
 */
async function expressRouteCrossConnectionBgpPeeringList() {
  const subscriptionId = "subid";
  const resourceGroupName = "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRouteCrossConnectionPeerings.list(
    resourceGroupName,
    crossConnectionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

expressRouteCrossConnectionBgpPeeringList().catch(console.error);
