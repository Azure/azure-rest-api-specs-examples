const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a ConfigurationPolicyGroup.
 *
 * @summary Deletes a ConfigurationPolicyGroup.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ConfigurationPolicyGroupDelete.json
 */
async function configurationPolicyGroupDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const vpnServerConfigurationName = "vpnServerConfiguration1";
  const configurationPolicyGroupName = "policyGroup1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.configurationPolicyGroups.beginDeleteAndWait(
    resourceGroupName,
    vpnServerConfigurationName,
    configurationPolicyGroupName
  );
  console.log(result);
}

configurationPolicyGroupDelete().catch(console.error);
