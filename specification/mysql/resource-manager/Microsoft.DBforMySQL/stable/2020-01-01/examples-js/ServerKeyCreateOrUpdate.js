const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a MySQL Server key.
 *
 * @summary Creates or updates a MySQL Server key.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerKeyCreateOrUpdate.json
 */
async function createsOrUpdatesAMySqlServerKey() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const serverName = "testserver";
  const keyName = "someVault_someKey_01234567890123456789012345678901";
  const resourceGroupName = "testrg";
  const parameters = {
    serverKeyType: "AzureKeyVault",
    uri: "https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901",
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.serverKeys.beginCreateOrUpdateAndWait(
    serverName,
    keyName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

createsOrUpdatesAMySqlServerKey().catch(console.error);
