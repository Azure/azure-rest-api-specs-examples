const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets Server Active Directory only authentication property or updates an existing server Active Directory only authentication property.
 *
 * @summary Sets Server Active Directory only authentication property or updates an existing server Active Directory only authentication property.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/AzureADOnlyAuthCreateOrUpdate.json
 */
async function createsOrUpdatesAzureActiveDirectoryOnlyAuthenticationObject() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-4799";
  const serverName = "sqlcrudtest-6440";
  const authenticationName = "Default";
  const parameters = {
    azureADOnlyAuthentication: false,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverAzureADOnlyAuthentications.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    authenticationName,
    parameters
  );
  console.log(result);
}

createsOrUpdatesAzureActiveDirectoryOnlyAuthenticationObject().catch(console.error);
