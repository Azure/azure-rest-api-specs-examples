const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function recommendedActionsListByServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const advisorName = "Index";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recommendedActions.listByServer(
    resourceGroupName,
    serverName,
    advisorName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

recommendedActionsListByServer().catch(console.error);
