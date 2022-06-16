const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a logical database's transparent data encryption.
 *
 * @summary Gets a logical database's transparent data encryption.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/TransparentDataEncryptionGet.json
 */
async function getADatabaseTransparentDataEncryption() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "security-tde-resourcegroup";
  const serverName = "securitytde";
  const databaseName = "testdb";
  const tdeName = "current";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.transparentDataEncryptions.get(
    resourceGroupName,
    serverName,
    databaseName,
    tdeName
  );
  console.log(result);
}

getADatabaseTransparentDataEncryption().catch(console.error);
