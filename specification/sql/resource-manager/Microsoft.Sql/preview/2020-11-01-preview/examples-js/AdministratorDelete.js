const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the Azure Active Directory administrator with the given name.
 *
 * @summary Deletes the Azure Active Directory administrator with the given name.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/AdministratorDelete.json
 */
async function deleteAzureActiveDirectoryAdministrator() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-4799";
  const serverName = "sqlcrudtest-6440";
  const administratorName = "ActiveDirectory";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverAzureADAdministrators.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    administratorName
  );
  console.log(result);
}

deleteAzureActiveDirectoryAdministrator().catch(console.error);
