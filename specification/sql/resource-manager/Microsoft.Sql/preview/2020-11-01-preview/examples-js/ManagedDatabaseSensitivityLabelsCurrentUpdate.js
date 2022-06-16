const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update sensitivity labels of a given database using an operations batch.
 *
 * @summary Update sensitivity labels of a given database using an operations batch.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsCurrentUpdate.json
 */
async function updateSensitivityLabelsOfAGivenDatabaseUsingAnOperationsBatch() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const managedInstanceName = "myManagedInstanceName";
  const databaseName = "myDatabase";
  const parameters = {
    operations: [
      {
        schema: "dbo",
        column: "column1",
        op: "set",
        sensitivityLabel: {
          informationType: "Financial",
          informationTypeId: "1D3652D6-422C-4115-82F1-65DAEBC665C8",
          labelId: "3A477B16-9423-432B-AA97-6069B481CEC3",
          labelName: "Highly Confidential",
        },
        table: "table1",
      },
      {
        schema: "dbo",
        column: "column2",
        op: "set",
        sensitivityLabel: {
          informationType: "PhoneNumber",
          informationTypeId: "d22fa6e9-5ee4-3bde-4c2b-a409604c4646",
          labelId: "bf91e08c-f4f0-478a-b016-25164b2a65ff",
          labelName: "PII",
        },
        table: "table2",
      },
      { schema: "dbo", column: "Column3", op: "remove", table: "Table1" },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedDatabaseSensitivityLabels.update(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    parameters
  );
  console.log(result);
}

updateSensitivityLabelsOfAGivenDatabaseUsingAnOperationsBatch().catch(console.error);
