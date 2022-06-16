const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of managed instances in a resource group.
 *
 * @summary Gets a list of managed instances in a resource group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceListByResourceGroupWithExpandEqualsAdministrators.json
 */
async function listManagedInstancesByResourceGroupWithExpandAdministratorsOrActivedirectory() {
  const subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
  const resourceGroupName = "Test1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedInstances.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedInstancesByResourceGroupWithExpandAdministratorsOrActivedirectory().catch(console.error);
