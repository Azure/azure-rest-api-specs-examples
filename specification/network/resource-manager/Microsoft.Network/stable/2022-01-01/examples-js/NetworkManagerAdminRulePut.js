const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an admin rule.
 *
 * @summary Creates or updates an admin rule.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerAdminRulePut.json
 */
async function createAnAdminRule() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestSecurityConfig";
  const ruleCollectionName = "testRuleCollection";
  const ruleName = "SampleAdminRule";
  const adminRule = {
    description: "This is Sample Admin Rule",
    access: "Deny",
    destinationPortRanges: ["22"],
    destinations: [{ addressPrefix: "*", addressPrefixType: "IPPrefix" }],
    direction: "Inbound",
    kind: "Custom",
    priority: 1,
    sourcePortRanges: ["0-65535"],
    sources: [{ addressPrefix: "Internet", addressPrefixType: "ServiceTag" }],
    protocol: "Tcp",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.adminRules.createOrUpdate(
    resourceGroupName,
    networkManagerName,
    configurationName,
    ruleCollectionName,
    ruleName,
    adminRule
  );
  console.log(result);
}

createAnAdminRule().catch(console.error);
