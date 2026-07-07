const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified nat gateway in a specified resource group.
 *
 * @summary gets the specified nat gateway in a specified resource group.
 * x-ms-original-file: 2025-07-01/NatGatewayWithNat64Get.json
 */
async function getNatGatewayWithNat64() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.natGateways.get("rg1", "test-natGateway");
  console.log(result);
}
