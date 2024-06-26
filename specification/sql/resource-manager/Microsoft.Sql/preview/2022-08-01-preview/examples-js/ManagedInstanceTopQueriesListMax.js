const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get top resource consuming queries of a managed instance.
 *
 * @summary Get top resource consuming queries of a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceTopQueriesListMax.json
 */
async function obtainListOfInstanceTopResourceConsumingQueriesFullBlownRequestAndResponse() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const databases = "db1,db2";
  const startTime = "2020-03-10T12:00:00Z";
  const endTime = "2020-03-12T12:00:00Z";
  const interval = "P1D";
  const observationMetric = "cpu";
  const options = {
    databases,
    startTime,
    endTime,
    interval,
    observationMetric,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedInstances.listByManagedInstance(
    resourceGroupName,
    managedInstanceName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
