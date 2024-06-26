const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get private DNS zone suffix in the cloud.
 *
 * @summary Get private DNS zone suffix in the cloud.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/stable/2023-12-30/examples/GetPrivateDnsZoneSuffix.json
 */
async function getPrivateDnsZoneSuffix() {
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential);
  const result = await client.getPrivateDnsZoneSuffix.execute();
  console.log(result);
}
