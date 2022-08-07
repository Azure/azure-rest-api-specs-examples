const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the rule collections in a security admin configuration, in a paginated format.
 *
 * @summary Lists all the rule collections in a security admin configuration, in a paginated format.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerAdminRuleCollectionList.json
 */
async function listSecurityAdminRuleCollections() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestSecurityConfig";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.adminRuleCollections.list(
    resourceGroupName,
    networkManagerName,
    configurationName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSecurityAdminRuleCollections().catch(console.error);
