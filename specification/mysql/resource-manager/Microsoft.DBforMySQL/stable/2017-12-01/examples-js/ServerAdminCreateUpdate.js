const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or update active directory administrator on an existing server. The update action will overwrite the existing administrator.
 *
 * @summary Creates or update active directory administrator on an existing server. The update action will overwrite the existing administrator.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerAdminCreateUpdate.json
 */
async function serverAdministratorCreate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mysqltestsvc4";
  const properties = {
    administratorType: "ActiveDirectory",
    login: "bob@contoso.com",
    sid: "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
    tenantId: "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.serverAdministrators.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    properties
  );
  console.log(result);
}

serverAdministratorCreate().catch(console.error);
