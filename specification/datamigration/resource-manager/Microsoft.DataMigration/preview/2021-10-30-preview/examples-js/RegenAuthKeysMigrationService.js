const { DataMigrationManagementClient } = require("@azure/arm-datamigration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate a new set of Authentication Keys for Self Hosted Integration Runtime.
 *
 * @summary Regenerate a new set of Authentication Keys for Self Hosted Integration Runtime.
 * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2021-10-30-preview/examples/RegenAuthKeysMigrationService.json
 */
async function regenerateTheOfAuthenticationKeys() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const sqlMigrationServiceName = "service1";
  const parameters = { keyName: "authKey1" };
  const credential = new DefaultAzureCredential();
  const client = new DataMigrationManagementClient(credential, subscriptionId);
  const result = await client.sqlMigrationServices.regenerateAuthKeys(
    resourceGroupName,
    sqlMigrationServiceName,
    parameters
  );
  console.log(result);
}

regenerateTheOfAuthenticationKeys().catch(console.error);
