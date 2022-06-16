const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function configurationCreateOrUpdate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "testserver";
  const configurationName = "event_scheduler";
  const parameters = { source: "user-override", value: "off" };
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.configurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    configurationName,
    parameters
  );
  console.log(result);
}

configurationCreateOrUpdate().catch(console.error);
