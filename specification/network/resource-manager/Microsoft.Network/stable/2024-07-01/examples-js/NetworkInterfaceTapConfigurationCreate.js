const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a Tap configuration in the specified NetworkInterface.
 *
 * @summary Creates or updates a Tap configuration in the specified NetworkInterface.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkInterfaceTapConfigurationCreate.json
 */
async function createNetworkInterfaceTapConfigurations() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "testrg";
  const networkInterfaceName = "mynic";
  const tapConfigurationName = "tapconfiguration1";
  const tapConfigurationParameters = {
    virtualNetworkTap: {
      id: "/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworkTaps/testvtap",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaceTapConfigurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkInterfaceName,
    tapConfigurationName,
    tapConfigurationParameters,
  );
  console.log(result);
}
