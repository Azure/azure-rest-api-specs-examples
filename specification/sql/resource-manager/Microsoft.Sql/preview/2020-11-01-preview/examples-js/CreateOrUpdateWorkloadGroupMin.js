const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a workload group.
 *
 * @summary Creates or updates a workload group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateWorkloadGroupMin.json
 */
async function createAWorkloadGroupWithTheRequiredPropertiesSpecified() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default-SQL-SouthEastAsia";
  const serverName = "testsvr";
  const databaseName = "testdb";
  const workloadGroupName = "smallrc";
  const parameters = {
    maxResourcePercent: 100,
    minResourcePercent: 0,
    minResourcePercentPerRequest: 3,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.workloadGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    workloadGroupName,
    parameters,
  );
  console.log(result);
}
