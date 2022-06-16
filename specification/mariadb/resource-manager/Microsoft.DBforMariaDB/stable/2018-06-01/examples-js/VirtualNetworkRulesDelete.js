const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteAVirtualNetworkRule() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "vnet-test-svr";
  const virtualNetworkRuleName = "vnet-firewall-rule";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkRules.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    virtualNetworkRuleName
  );
  console.log(result);
}

deleteAVirtualNetworkRule().catch(console.error);
