const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an express route circuit.
 *
 * @summary creates or updates an express route circuit.
 * x-ms-original-file: 2025-07-01/ExpressRouteCircuitCreateOnExpressRoutePort.json
 */
async function createExpressRouteCircuitOnExpressRoutePort() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.createOrUpdate("rg1", "expressRouteCircuit1", {
    location: "westus",
    authorizationKey: "b0be57f5-1fba-463b-adec-ffe767354cdd",
    bandwidthInGbps: 10,
    enableDirectPortRateLimit: false,
    expressRoutePort: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName",
    },
    sku: { name: "Premium_MeteredData", family: "MeteredData", tier: "Premium" },
  });
  console.log(result);
}
