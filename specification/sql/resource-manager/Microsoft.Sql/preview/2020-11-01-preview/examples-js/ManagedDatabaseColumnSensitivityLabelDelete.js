const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the sensitivity label of a given column
 *
 * @summary Deletes the sensitivity label of a given column
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnSensitivityLabelDelete.json
 */
async function deletesTheSensitivityLabelOfAGivenColumnInAManagedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const schemaName = "dbo";
  const tableName = "myTable";
  const columnName = "myColumn";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseSensitivityLabels.delete(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    schemaName,
    tableName,
    columnName
  );
  console.log(result);
}

deletesTheSensitivityLabelOfAGivenColumnInAManagedDatabase().catch(console.error);
