const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all managed instances in an instance pool.
 *
 * @summary Gets a list of all managed instances in an instance pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceListByInstancePoolWithExpandEqualsAdministrators.json
 */
async function listManagedInstancesByInstancePoolWithExpandAdministratorsOrActivedirectory() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "20D7082A-0FC7-4468-82BD-542694D5042B";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Test1";
  const instancePoolName = "pool1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedInstances.listByInstancePool(
    resourceGroupName,
    instancePoolName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
