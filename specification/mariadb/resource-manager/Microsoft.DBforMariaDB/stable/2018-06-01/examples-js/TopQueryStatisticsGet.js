const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function topQueryStatisticsGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const queryStatisticId = "66-636923268000000000-636923277000000000-avg-duration";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.topQueryStatistics.get(
    resourceGroupName,
    serverName,
    queryStatisticId
  );
  console.log(result);
}

topQueryStatisticsGet().catch(console.error);
