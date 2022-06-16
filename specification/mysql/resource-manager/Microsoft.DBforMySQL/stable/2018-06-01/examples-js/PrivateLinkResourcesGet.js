const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a private link resource for MySQL server.
 *
 * @summary Gets a private link resource for MySQL server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateLinkResourcesGet.json
 */
async function getsAPrivateLinkResourceForMySql() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "test-svr";
  const groupName = "plr";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, serverName, groupName);
  console.log(result);
}

getsAPrivateLinkResourceForMySql().catch(console.error);
