const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function waitStatisticsGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const waitStatisticsId =
    "636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.waitStatistics.get(resourceGroupName, serverName, waitStatisticsId);
  console.log(result);
}

waitStatisticsGet().catch(console.error);
