const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a nat rule to a scalable vpn gateway if it doesn't exist else updates the existing nat rules.
 *
 * @summary Creates a nat rule to a scalable vpn gateway if it doesn't exist else updates the existing nat rules.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NatRulePut.json
 */
async function natRulePut() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const gatewayName = "gateway1";
  const natRuleName = "natRule1";
  const natRuleParameters = {
    typePropertiesType: "Static",
    externalMappings: [{ addressSpace: "192.168.21.0/24" }],
    internalMappings: [{ addressSpace: "10.4.0.0/24" }],
    ipConfigurationId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/cloudnet1-VNG/ipConfigurations/default",
    mode: "EgressSnat",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.natRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    gatewayName,
    natRuleName,
    natRuleParameters,
  );
  console.log(result);
}
