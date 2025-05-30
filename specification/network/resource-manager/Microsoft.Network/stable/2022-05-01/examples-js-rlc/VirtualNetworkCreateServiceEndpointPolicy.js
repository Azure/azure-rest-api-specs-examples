const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a virtual network in the specified resource group.
 *
 * @summary Creates or updates a virtual network in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkCreateServiceEndpointPolicy.json
 */
async function createVirtualNetworkWithServiceEndpointsAndServiceEndpointPolicy() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "vnetTest";
  const virtualNetworkName = "vnet1";
  const options = {
    body: {
      location: "eastus2euap",
      properties: {
        addressSpace: { addressPrefixes: ["10.0.0.0/16"] },
        subnets: [
          {
            name: "test-1",
            properties: {
              addressPrefix: "10.0.0.0/16",
              serviceEndpointPolicies: [
                {
                  id: "/subscriptions/subid/resourceGroups/vnetTest/providers/Microsoft.Network/serviceEndpointPolicies/ServiceEndpointPolicy1",
                },
              ],
              serviceEndpoints: [{ service: "Microsoft.Storage" }],
            },
          },
        ],
      },
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}",
      subscriptionId,
      resourceGroupName,
      virtualNetworkName,
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createVirtualNetworkWithServiceEndpointsAndServiceEndpointPolicy().catch(console.error);
