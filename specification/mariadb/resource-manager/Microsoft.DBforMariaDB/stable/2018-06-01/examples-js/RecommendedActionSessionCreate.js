const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function recommendedActionSessionCreate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const advisorName = "Index";
  const databaseName = "someDatabaseName";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.beginCreateRecommendedActionSessionAndWait(
    resourceGroupName,
    serverName,
    advisorName,
    databaseName
  );
  console.log(result);
}

recommendedActionSessionCreate().catch(console.error);
