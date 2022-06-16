const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an express route circuit.
 *
 * @summary Creates or updates an express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitCreateOnExpressRoutePort.json
 */
async function createExpressRouteCircuitOnExpressRoutePort() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "expressRouteCircuit1";
  const parameters = {
    bandwidthInGbps: 10,
    expressRoutePort: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName",
    },
    location: "westus",
    sku: { name: "Premium_MeteredData", family: "MeteredData", tier: "Premium" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.beginCreateOrUpdateAndWait(
    resourceGroupName,
    circuitName,
    parameters
  );
  console.log(result);
}

createExpressRouteCircuitOnExpressRoutePort().catch(console.error);
