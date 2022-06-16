const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of the logical database's transparent data encryption.
 *
 * @summary Gets a list of the logical database's transparent data encryption.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/TransparentDataEncryptionList.json
 */
async function getAListOfTheDatabaseTransparentDataEncryption() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "security-tde-resourcegroup";
  const serverName = "securitytde";
  const databaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.transparentDataEncryptions.listByDatabase(
    resourceGroupName,
    serverName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAListOfTheDatabaseTransparentDataEncryption().catch(console.error);
