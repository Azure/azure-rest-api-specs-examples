const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates an express route circuit.
 *
 * @summary Creates or updates an express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteCircuitCreateOnExpressRoutePort.json
 */
async function createExpressRouteCircuitOnExpressRoutePort() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const circuitName = "expressRouteCircuit1";
  const options = {
    body: {
      location: "westus",
      properties: {
        authorizationKey: "b0be57f5-1fba-463b-adec-ffe767354cdd",
        bandwidthInGbps: 10,
        expressRoutePort: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName",
        },
      },
      sku: {
        name: "Premium_MeteredData",
        family: "MeteredData",
        tier: "Premium",
      },
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}",
      subscriptionId,
      resourceGroupName,
      circuitName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createExpressRouteCircuitOnExpressRoutePort().catch(console.error);
