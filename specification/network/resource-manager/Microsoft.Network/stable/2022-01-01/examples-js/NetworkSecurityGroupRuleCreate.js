const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a security rule in the specified network security group.
 *
 * @summary Creates or updates a security rule in the specified network security group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkSecurityGroupRuleCreate.json
 */
async function createSecurityRule() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkSecurityGroupName = "testnsg";
  const securityRuleName = "rule1";
  const securityRuleParameters = {
    access: "Deny",
    destinationAddressPrefix: "11.0.0.0/8",
    destinationPortRange: "8080",
    direction: "Outbound",
    priority: 100,
    sourceAddressPrefix: "10.0.0.0/8",
    sourcePortRange: "*",
    protocol: "*",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.securityRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkSecurityGroupName,
    securityRuleName,
    securityRuleParameters
  );
  console.log(result);
}

createSecurityRule().catch(console.error);
