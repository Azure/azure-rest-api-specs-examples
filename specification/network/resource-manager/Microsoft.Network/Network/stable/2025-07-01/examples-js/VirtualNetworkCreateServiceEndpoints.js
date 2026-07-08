const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a virtual network in the specified resource group.
 *
 * @summary creates or updates a virtual network in the specified resource group.
 * x-ms-original-file: 2025-07-01/VirtualNetworkCreateServiceEndpoints.json
 */
async function createVirtualNetworkWithServiceEndpoints() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworks.createOrUpdate("vnetTest", "vnet1", {
    location: "eastus",
    addressSpace: { addressPrefixes: ["10.0.0.0/16"] },
    subnets: [
      { addressPrefix: "10.0.0.0/16", serviceEndpoints: [{ service: "Microsoft.Storage" }] },
    ],
  });
  console.log(result);
}
