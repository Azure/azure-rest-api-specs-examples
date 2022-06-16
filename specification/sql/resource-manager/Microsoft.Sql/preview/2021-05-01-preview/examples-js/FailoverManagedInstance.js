const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Failovers a managed instance.
 *
 * @summary Failovers a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/FailoverManagedInstance.json
 */
async function failoverAManagedInstance() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const managedInstanceName = "instanceName";
  const replicaType = "Primary";
  const options = { replicaType };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstances.beginFailoverAndWait(
    resourceGroupName,
    managedInstanceName,
    options
  );
  console.log(result);
}

failoverAManagedInstance().catch(console.error);
