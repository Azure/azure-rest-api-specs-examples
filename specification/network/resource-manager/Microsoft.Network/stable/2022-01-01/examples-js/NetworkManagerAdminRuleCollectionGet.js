const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a network manager security admin configuration rule collection.
 *
 * @summary Gets a network manager security admin configuration rule collection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerAdminRuleCollectionGet.json
 */
async function getsSecurityAdminRuleCollection() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestSecurityConfig";
  const ruleCollectionName = "testRuleCollection";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.adminRuleCollections.get(
    resourceGroupName,
    networkManagerName,
    configurationName,
    ruleCollectionName
  );
  console.log(result);
}

getsSecurityAdminRuleCollection().catch(console.error);
