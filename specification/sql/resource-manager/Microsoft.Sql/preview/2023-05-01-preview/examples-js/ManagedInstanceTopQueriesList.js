const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get top resource consuming queries of a managed instance.
 *
 * @summary Get top resource consuming queries of a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ManagedInstanceTopQueriesList.json
 */
async function obtainListOfInstanceTopResourceConsumingQueries() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const interval = "PT1H";
  const observationMetric = "duration";
  const options = {
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
