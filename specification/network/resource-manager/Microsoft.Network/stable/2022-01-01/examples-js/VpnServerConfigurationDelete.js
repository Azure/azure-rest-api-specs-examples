const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a VpnServerConfiguration.
 *
 * @summary Deletes a VpnServerConfiguration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnServerConfigurationDelete.json
 */
async function vpnServerConfigurationDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const vpnServerConfigurationName = "vpnServerConfiguration1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnServerConfigurations.beginDeleteAndWait(
    resourceGroupName,
    vpnServerConfigurationName
  );
  console.log(result);
}

vpnServerConfigurationDelete().catch(console.error);
