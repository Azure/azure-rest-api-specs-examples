const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the MySQL Server key with the given name.
 *
 * @summary Deletes the MySQL Server key with the given name.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerKeyDelete.json
 */
async function deleteTheMySqlServerKey() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const serverName = "testserver";
  const keyName = "someVault_someKey_01234567890123456789012345678901";
  const resourceGroupName = "testrg";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.serverKeys.beginDeleteAndWait(serverName, keyName, resourceGroupName);
  console.log(result);
}

deleteTheMySqlServerKey().catch(console.error);
