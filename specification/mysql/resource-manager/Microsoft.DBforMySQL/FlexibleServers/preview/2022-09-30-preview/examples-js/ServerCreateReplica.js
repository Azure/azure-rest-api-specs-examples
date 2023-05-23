const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary Creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2022-09-30-preview/examples/ServerCreateReplica.json
 */
async function createAReplicaServer() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["MYSQL_RESOURCE_GROUP"] || "testgr";
  const serverName = "replica-server";
  const parameters = {
    createMode: "Replica",
    location: "SoutheastAsia",
    sourceServerResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testgr/providers/Microsoft.DBforMySQL/flexibleServers/source-server",
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginCreateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}
