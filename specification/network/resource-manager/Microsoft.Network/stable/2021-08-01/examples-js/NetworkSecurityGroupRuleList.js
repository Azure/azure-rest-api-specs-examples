const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all security rules in a network security group.
 *
 * @summary Gets all security rules in a network security group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkSecurityGroupRuleList.json
 */
async function listNetworkSecurityRulesInNetworkSecurityGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkSecurityGroupName = "testnsg";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.securityRules.list(resourceGroupName, networkSecurityGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listNetworkSecurityRulesInNetworkSecurityGroup().catch(console.error);
