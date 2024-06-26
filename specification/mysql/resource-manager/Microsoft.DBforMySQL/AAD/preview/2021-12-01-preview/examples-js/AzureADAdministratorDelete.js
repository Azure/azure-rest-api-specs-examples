const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an Azure AD Administrator.
 *
 * @summary Deletes an Azure AD Administrator.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/preview/2021-12-01-preview/examples/AzureADAdministratorDelete.json
 */
async function deleteAnAzureAdAdministrator() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["MYSQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "mysqltestsvc4";
  const administratorName = "ActiveDirectory";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.azureADAdministrators.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    administratorName
  );
  console.log(result);
}
