const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the sensitivity label of a given column
 *
 * @summary Gets the sensitivity label of a given column
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseColumnSensitivityLabelGet.json
 */
async function getsTheSensitivityLabelOfAGivenColumnInAManagedDatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const schemaName = "dbo";
  const tableName = "myTable";
  const columnName = "myColumn";
  const sensitivityLabelSource = "current";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseSensitivityLabels.get(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    schemaName,
    tableName,
    columnName,
    sensitivityLabelSource
  );
  console.log(result);
}

getsTheSensitivityLabelOfAGivenColumnInAManagedDatabase().catch(console.error);
