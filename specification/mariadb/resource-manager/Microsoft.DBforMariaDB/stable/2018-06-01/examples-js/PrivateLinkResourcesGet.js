const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function getsAPrivateLinkResourceForMariaDb() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "test-svr";
  const groupName = "plr";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, serverName, groupName);
  console.log(result);
}

getsAPrivateLinkResourceForMariaDb().catch(console.error);
