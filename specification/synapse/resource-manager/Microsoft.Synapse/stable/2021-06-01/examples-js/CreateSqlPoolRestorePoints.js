const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a restore point for a data warehouse.
 *
 * @summary Creates a restore point for a data warehouse.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPoolRestorePoints.json
 */
async function createsSqlPoolRestorePoint() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default-SQL-SouthEastAsia";
  const workspaceName = "testserver";
  const sqlPoolName = "testDatabase";
  const parameters = {
    restorePointLabel: "mylabel",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.sqlPoolRestorePoints.beginCreateAndWait(
    resourceGroupName,
    workspaceName,
    sqlPoolName,
    parameters
  );
  console.log(result);
}

createsSqlPoolRestorePoint().catch(console.error);
