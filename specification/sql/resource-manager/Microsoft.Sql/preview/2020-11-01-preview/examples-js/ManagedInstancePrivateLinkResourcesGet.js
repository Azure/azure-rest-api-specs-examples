const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a private link resource for SQL server.
 *
 * @summary Gets a private link resource for SQL server.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstancePrivateLinkResourcesGet.json
 */
async function getsAPrivateLinkResourceForSql() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const managedInstanceName = "test-cl";
  const groupName = "plr";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstancePrivateLinkResources.get(
    resourceGroupName,
    managedInstanceName,
    groupName
  );
  console.log(result);
}

getsAPrivateLinkResourceForSql().catch(console.error);
