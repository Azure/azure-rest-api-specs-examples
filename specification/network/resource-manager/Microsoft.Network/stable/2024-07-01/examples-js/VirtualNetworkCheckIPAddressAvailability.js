const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Checks whether a private IP address is available for use.
 *
 * @summary Checks whether a private IP address is available for use.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualNetworkCheckIPAddressAvailability.json
 */
async function checkIPAddressAvailability() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkName = "test-vnet";
  const ipAddress = "10.0.1.4";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworks.checkIPAddressAvailability(
    resourceGroupName,
    virtualNetworkName,
    ipAddress,
  );
  console.log(result);
}
