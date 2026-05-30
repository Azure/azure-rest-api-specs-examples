const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: 2025-06-01-preview/ServerCreateReplica.json
 */
async function createAReplicaServer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.create("testgr", "replica-server", {
    location: "SoutheastAsia",
    createMode: "Replica",
    sourceServerResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testgr/providers/Microsoft.DBforMySQL/flexibleServers/source-server",
  });
  console.log(result);
}
