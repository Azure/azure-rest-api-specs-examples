const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a managed instance.
 *
 * @summary Updates a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceUpdateMin.json
 */
async function updateManagedInstanceWithMinimalProperties() {
  const subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
  const resourceGroupName = "testrg";
  const managedInstanceName = "testinstance";
  const parameters = { tags: { tagKey1: "TagValue1" } };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstances.beginUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    parameters
  );
  console.log(result);
}

updateManagedInstanceWithMinimalProperties().catch(console.error);
