const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of managed database's transparent data encryptions.
 *
 * @summary Gets a list of managed database's transparent data encryptions.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedTransparentDataEncryptionList.json
 */
async function getAListOfTheDatabaseTransparentDataEncryptions() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "security-tde-resourcegroup";
  const managedInstanceName = "securitytde";
  const databaseName = "testdb";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseTransparentDataEncryption.listByDatabase(
    resourceGroupName,
    managedInstanceName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAListOfTheDatabaseTransparentDataEncryptions().catch(console.error);
