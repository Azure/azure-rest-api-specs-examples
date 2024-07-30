const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of a VpnServerConfiguration.
 *
 * @summary Retrieves the details of a VpnServerConfiguration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/VpnServerConfigurationGet.json
 */
async function vpnServerConfigurationGet() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const vpnServerConfigurationName = "vpnServerConfiguration1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnServerConfigurations.get(
    resourceGroupName,
    vpnServerConfigurationName,
  );
  console.log(result);
}
