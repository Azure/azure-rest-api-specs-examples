const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a specific Azure Active Directory only authentication property.
 *
 * @summary Gets a specific Azure Active Directory only authentication property.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceAzureADOnlyAuthGet.json
 */
async function getsAAzureActiveDirectoryOnlyAuthenticationProperty() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const managedInstanceName = "managedInstance";
  const authenticationName = "Default";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceAzureADOnlyAuthentications.get(
    resourceGroupName,
    managedInstanceName,
    authenticationName,
  );
  console.log(result);
}
