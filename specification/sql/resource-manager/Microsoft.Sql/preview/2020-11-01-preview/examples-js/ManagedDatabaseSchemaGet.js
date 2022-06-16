const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get managed database schema
 *
 * @summary Get managed database schema
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSchemaGet.json
 */
async function getManagedDatabaseSchema() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const schemaName = "dbo";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseSchemas.get(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    schemaName
  );
  console.log(result);
}

getManagedDatabaseSchema().catch(console.error);
