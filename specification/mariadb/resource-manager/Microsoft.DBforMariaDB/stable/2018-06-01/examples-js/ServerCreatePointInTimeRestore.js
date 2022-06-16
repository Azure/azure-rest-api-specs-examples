const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function createADatabaseAsAPointInTimeRestore() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TargetResourceGroup";
  const serverName = "targetserver";
  const parameters = {
    location: "brazilsouth",
    properties: {
      createMode: "PointInTimeRestore",
      restorePointInTime: new Date("2017-12-14T00:00:37.467Z"),
      sourceServerId:
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMariaDB/servers/sourceserver",
    },
    sku: {
      name: "GP_Gen5_2",
      capacity: 2,
      family: "Gen5",
      tier: "GeneralPurpose",
    },
    tags: { elasticServer: "1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

createADatabaseAsAPointInTimeRestore().catch(console.error);
