const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get private DNS zone suffix in the cloud.
 *
 * @summary Get private DNS zone suffix in the cloud.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2021-12-01-preview/examples/GetPrivateDnsZoneSuffix.json
 */
async function getPrivateDnsZoneSuffix() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.getPrivateDnsZoneSuffix.execute();
  console.log(result);
}
