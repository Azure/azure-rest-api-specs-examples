const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disables sensitivity recommendations on a given column
 *
 * @summary Disables sensitivity recommendations on a given column
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/RecommendedColumnSensitivityLabelDisable.json
 */
async function disablesSensitivityRecommendationsOnAGivenColumn() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const serverName = "myServer";
  const databaseName = "myDatabase";
  const schemaName = "dbo";
  const tableName = "myTable";
  const columnName = "myColumn";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.sensitivityLabels.disableRecommendation(
    resourceGroupName,
    serverName,
    databaseName,
    schemaName,
    tableName,
    columnName
  );
  console.log(result);
}

disablesSensitivityRecommendationsOnAGivenColumn().catch(console.error);
