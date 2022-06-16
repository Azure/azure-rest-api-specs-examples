const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all managed instances in the subscription.
 *
 * @summary Gets a list of all managed instances in the subscription.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceListWithExpandEqualsAdministrators.json
 */
async function listManagedInstancesWithExpandAdministratorsOrActivedirectory() {
  const subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedInstances.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedInstancesWithExpandAdministratorsOrActivedirectory().catch(console.error);
