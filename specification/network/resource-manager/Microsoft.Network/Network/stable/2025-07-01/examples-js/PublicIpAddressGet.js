const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified public IP address in a specified resource group.
 *
 * @summary gets the specified public IP address in a specified resource group.
 * x-ms-original-file: 2025-07-01/PublicIpAddressGet.json
 */
async function getPublicIPAddress() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPAddresses.get("rg1", "testDNS-ip");
  console.log(result);
}
