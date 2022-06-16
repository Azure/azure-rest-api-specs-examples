const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of a VpnServerConfiguration.
 *
 * @summary Retrieves the details of a VpnServerConfiguration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnServerConfigurationGet.json
 */
async function vpnServerConfigurationGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const vpnServerConfigurationName = "vpnServerConfiguration1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnServerConfigurations.get(
    resourceGroupName,
    vpnServerConfigurationName
  );
  console.log(result);
}

vpnServerConfigurationGet().catch(console.error);
