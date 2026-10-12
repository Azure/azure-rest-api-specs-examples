const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a capability on a virtual network appliance.
 *
 * @summary creates or updates a capability on a virtual network appliance.
 * x-ms-original-file: 2026-03-01/VirtualNetworkApplianceCapabilities_CreateOrUpdate_PLGateway.json
 */
async function createVirtualNetworkAppliancePLGatewayCapability() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkApplianceCapabilities.createOrUpdate(
    "rg1",
    "test-vna",
    "pl-gateway",
    { kind: "PLGateway", properties: { ipVersion: "IPv6" } },
  );
  console.log(result);
}
