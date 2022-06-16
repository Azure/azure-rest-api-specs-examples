const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resets the VpnLink connection specified.
 *
 * @summary Resets the VpnLink connection specified.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnSiteLinkConnectionReset.json
 */
async function resetVpnLinkConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const connectionName = "vpnConnection1";
  const linkConnectionName = "Connection-Link1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnLinkConnections.beginResetConnectionAndWait(
    resourceGroupName,
    gatewayName,
    connectionName,
    linkConnectionName
  );
  console.log(result);
}

resetVpnLinkConnection().catch(console.error);
