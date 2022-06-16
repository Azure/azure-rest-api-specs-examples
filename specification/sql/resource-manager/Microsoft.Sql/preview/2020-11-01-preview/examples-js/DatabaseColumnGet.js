const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get database column
 *
 * @summary Get database column
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseColumnGet.json
 */
async function getDatabaseColumn() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "myRG";
  const serverName = "serverName";
  const databaseName = "myDatabase";
  const schemaName = "dbo";
  const tableName = "table1";
  const columnName = "column1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseColumns.get(
    resourceGroupName,
    serverName,
    databaseName,
    schemaName,
    tableName,
    columnName
  );
  console.log(result);
}

getDatabaseColumn().catch(console.error);
