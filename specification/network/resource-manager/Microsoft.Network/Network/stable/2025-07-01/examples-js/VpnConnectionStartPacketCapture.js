const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to starts packet capture on Vpn connection in the specified resource group.
 *
 * @summary starts packet capture on Vpn connection in the specified resource group.
 * x-ms-original-file: 2025-07-01/VpnConnectionStartPacketCapture.json
 */
async function startPacketCaptureOnVpnConnectionWithoutFilter() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnConnections.startPacketCapture(
    "rg1",
    "gateway1",
    "vpnConnection1",
    { parameters: { linkConnectionNames: ["siteLink1", "siteLink2"] } },
  );
  console.log(result);
}
