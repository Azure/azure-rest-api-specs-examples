const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists IKE Security Associations for Vpn Site Link Connection in the specified resource group.
 *
 * @summary Lists IKE Security Associations for Vpn Site Link Connection in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VpnSiteLinkConnectionGetIkeSas.json
 */
async function getVpnLinkConnectionIkeSa() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const connectionName = "vpnConnection1";
  const linkConnectionName = "Connection-Link1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnLinkConnections.beginGetIkeSasAndWait(
    resourceGroupName,
    gatewayName,
    connectionName,
    linkConnectionName
  );
  console.log(result);
}

getVpnLinkConnectionIkeSa().catch(console.error);
