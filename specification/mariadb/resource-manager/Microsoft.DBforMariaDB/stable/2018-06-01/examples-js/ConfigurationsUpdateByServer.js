const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function configurationList() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mariadbtestsvc1";
  const value = {};
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.serverParameters.beginListUpdateConfigurationsAndWait(
    resourceGroupName,
    serverName,
    value
  );
  console.log(result);
}

configurationList().catch(console.error);
