const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a network manager security admin configuration.
 *
 * @summary creates or updates a network manager security admin configuration.
 * x-ms-original-file: 2025-07-01/NetworkManagerSecurityAdminConfigurationPut.json
 */
async function createNetworkManagerSecurityAdminConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.securityAdminConfigurations.createOrUpdate(
    "rg1",
    "testNetworkManager",
    "myTestSecurityConfig",
    { description: "A sample policy", applyOnNetworkIntentPolicyBasedServices: ["None"] },
  );
  console.log(result);
}
