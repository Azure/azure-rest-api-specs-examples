const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
 *
 * @summary Gives the list of VpnServerConfigurations associated with Virtual Wan in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/GetVirtualWanVpnServerConfigurations.json
 */
async function getVirtualWanVpnServerConfigurations() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualWANName = "wan1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnServerConfigurationsAssociatedWithVirtualWan.beginListAndWait(
    resourceGroupName,
    virtualWANName
  );
  console.log(result);
}

getVirtualWanVpnServerConfigurations().catch(console.error);
