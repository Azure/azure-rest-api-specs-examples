const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of a vpn site link connection.
 *
 * @summary Retrieves the details of a vpn site link connection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnSiteLinkConnectionGet.json
 */
async function vpnSiteLinkConnectionGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const connectionName = "vpnConnection1";
  const linkConnectionName = "Connection-Link1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnSiteLinkConnections.get(
    resourceGroupName,
    gatewayName,
    connectionName,
    linkConnectionName
  );
  console.log(result);
}

vpnSiteLinkConnectionGet().catch(console.error);
