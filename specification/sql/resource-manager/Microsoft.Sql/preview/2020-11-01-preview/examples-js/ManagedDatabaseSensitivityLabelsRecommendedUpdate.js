const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update recommended sensitivity labels states of a given database using an operations batch.
 *
 * @summary Update recommended sensitivity labels states of a given database using an operations batch.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsRecommendedUpdate.json
 */
async function updateRecommendedSensitivityLabelsOfAGivenDatabaseUsingAnOperationsBatch() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const parameters = {
    operations: [
      { schema: "dbo", column: "column1", op: "enable", table: "table1" },
      { schema: "dbo", column: "column2", op: "disable", table: "table2" },
      { schema: "dbo", column: "Column3", op: "disable", table: "Table1" },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseRecommendedSensitivityLabels.update(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    parameters,
  );
  console.log(result);
}
