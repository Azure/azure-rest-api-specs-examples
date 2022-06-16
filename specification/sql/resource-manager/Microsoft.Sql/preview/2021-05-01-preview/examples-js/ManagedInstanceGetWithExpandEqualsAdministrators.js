const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a managed instance.
 *
 * @summary Gets a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceGetWithExpandEqualsAdministrators.json
 */
async function getManagedInstanceWithExpandAdministratorsOrActivedirectory() {
  const subscriptionId = "20d7082a-0fc7-4468-82bd-542694d5042b";
  const resourceGroupName = "testrg";
  const managedInstanceName = "testinstance";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstances.get(resourceGroupName, managedInstanceName);
  console.log(result);
}

getManagedInstanceWithExpandAdministratorsOrActivedirectory().catch(console.error);
