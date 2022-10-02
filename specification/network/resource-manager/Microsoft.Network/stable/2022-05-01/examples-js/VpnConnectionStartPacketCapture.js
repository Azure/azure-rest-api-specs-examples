const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts packet capture on Vpn connection in the specified resource group.
 *
 * @summary Starts packet capture on Vpn connection in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VpnConnectionStartPacketCapture.json
 */
async function startPacketCaptureOnVpnConnectionWithoutFilter() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const vpnConnectionName = "vpnConnection1";
  const parameters = {
    linkConnectionNames: ["siteLink1", "siteLink2"],
  };
  const options = {
    parameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnConnections.beginStartPacketCaptureAndWait(
    resourceGroupName,
    gatewayName,
    vpnConnectionName,
    options
  );
  console.log(result);
}

startPacketCaptureOnVpnConnectionWithoutFilter().catch(console.error);
