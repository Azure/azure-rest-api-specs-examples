const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the currently advertised routes table summary associated with the express route circuit in a resource group.
 *
 * @summary Gets the currently advertised routes table summary associated with the express route circuit in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitRouteTableSummaryList.json
 */
async function listRouteTableSummary() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "circuitName";
  const peeringName = "peeringName";
  const devicePath = "devicePath";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.beginListRoutesTableSummaryAndWait(
    resourceGroupName,
    circuitName,
    peeringName,
    devicePath
  );
  console.log(result);
}

listRouteTableSummary().catch(console.error);
