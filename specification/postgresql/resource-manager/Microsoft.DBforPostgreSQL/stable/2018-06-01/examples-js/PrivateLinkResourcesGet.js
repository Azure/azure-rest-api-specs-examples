const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a private link resource for PostgreSQL server.
 *
 * @summary Gets a private link resource for PostgreSQL server.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateLinkResourcesGet.json
 */
async function getsAPrivateLinkResourceForPostgreSql() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "test-svr";
  const groupName = "plr";
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, serverName, groupName);
  console.log(result);
}

getsAPrivateLinkResourceForPostgreSql().catch(console.error);
