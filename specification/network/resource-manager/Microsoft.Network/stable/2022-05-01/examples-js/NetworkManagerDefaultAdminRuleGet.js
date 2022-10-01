const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a network manager security configuration admin rule.
 *
 * @summary Gets a network manager security configuration admin rule.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerDefaultAdminRuleGet.json
 */
async function getsSecurityDefaultAdminRule() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestSecurityConfig";
  const ruleCollectionName = "testRuleCollection";
  const ruleName = "SampleDefaultAdminRule";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.adminRules.get(
    resourceGroupName,
    networkManagerName,
    configurationName,
    ruleCollectionName,
    ruleName
  );
  console.log(result);
}

getsSecurityDefaultAdminRule().catch(console.error);
