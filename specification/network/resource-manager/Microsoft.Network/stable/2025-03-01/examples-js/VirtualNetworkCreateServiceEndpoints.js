const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a virtual network in the specified resource group.
 *
 * @summary Creates or updates a virtual network in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkCreateServiceEndpoints.json
 */
async function createVirtualNetworkWithServiceEndpoints() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "vnetTest";
  const virtualNetworkName = "vnet1";
  const parameters = {
    addressSpace: { addressPrefixes: ["10.0.0.0/16"] },
    location: "eastus",
    subnets: [
      {
        name: "test-1",
        addressPrefix: "10.0.0.0/16",
        serviceEndpoints: [{ service: "Microsoft.Storage" }],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualNetworkName,
    parameters,
  );
  console.log(result);
}
