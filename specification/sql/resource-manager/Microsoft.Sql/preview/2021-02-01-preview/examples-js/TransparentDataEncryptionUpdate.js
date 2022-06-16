const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a logical database's transparent data encryption configuration.
 *
 * @summary Updates a logical database's transparent data encryption configuration.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/TransparentDataEncryptionUpdate.json
 */
async function updateADatabaseTransparentDataEncryptionStateWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securitytde-42-rg";
  const serverName = "securitytde-42";
  const databaseName = "testdb";
  const tdeName = "current";
  const parameters = {
    state: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.transparentDataEncryptions.createOrUpdate(
    resourceGroupName,
    serverName,
    databaseName,
    tdeName,
    parameters
  );
  console.log(result);
}

updateADatabaseTransparentDataEncryptionStateWithMinimalParameters().catch(console.error);
