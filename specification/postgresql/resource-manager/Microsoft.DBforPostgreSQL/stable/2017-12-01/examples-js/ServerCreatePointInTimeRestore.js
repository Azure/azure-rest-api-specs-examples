const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new server, or will overwrite an existing server.
 *
 * @summary Creates a new server, or will overwrite an existing server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerCreatePointInTimeRestore.json
 */
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
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforPostgreSQL/servers/sourceserver",
    },
    sku: { name: "B_Gen5_2", capacity: 2, family: "Gen5", tier: "Basic" },
    tags: { elasticServer: "1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}

createADatabaseAsAPointInTimeRestore().catch(console.error);
