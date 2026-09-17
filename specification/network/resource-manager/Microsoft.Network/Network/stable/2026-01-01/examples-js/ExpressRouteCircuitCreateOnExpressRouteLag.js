const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an express route circuit.
 *
 * @summary creates or updates an express route circuit.
 * x-ms-original-file: 2026-01-01/ExpressRouteCircuitCreateOnExpressRouteLag.json
 */
async function createExpressRouteCircuitOnExpressRouteLag() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.createOrUpdate("rg1", "expressRouteCircuit1", {
    location: "eastus2euap",
    bandwidthInGbps: 5,
    enableDirectPortRateLimit: true,
    expressRouteLag: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteLags/lagName",
    },
    sku: { name: "Premium_MeteredData", family: "MeteredData", tier: "Premium" },
  });
  console.log(result);
}
