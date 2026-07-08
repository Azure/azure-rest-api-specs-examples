const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an express route circuit.
 *
 * @summary creates or updates an express route circuit.
 * x-ms-original-file: 2025-07-01/ExpressRouteCircuitCreate.json
 */
async function createExpressRouteCircuit() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.createOrUpdate("rg1", "circuitName", {
    location: "Brazil South",
    allowClassicOperations: false,
    authorizations: [],
    peerings: [],
    serviceProviderProperties: {
      bandwidthInMbps: 200,
      peeringLocation: "Silicon Valley",
      serviceProviderName: "Equinix",
    },
    sku: { name: "Standard_MeteredData", family: "MeteredData", tier: "Standard" },
  });
  console.log(result);
}
