const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an express route circuit.
 *
 * @summary Creates or updates an express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitCreate.json
 */
async function createExpressRouteCircuit() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "circuitName";
  const parameters = {
    allowClassicOperations: false,
    authorizations: [],
    location: "Brazil South",
    peerings: [],
    serviceProviderProperties: {
      bandwidthInMbps: 200,
      peeringLocation: "Silicon Valley",
      serviceProviderName: "Equinix",
    },
    sku: {
      name: "Standard_MeteredData",
      family: "MeteredData",
      tier: "Standard",
    },
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

createExpressRouteCircuit().catch(console.error);
