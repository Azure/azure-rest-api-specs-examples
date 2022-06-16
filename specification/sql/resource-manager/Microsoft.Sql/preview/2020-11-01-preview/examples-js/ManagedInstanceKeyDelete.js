const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the managed instance key with the given name.
 *
 * @summary Deletes the managed instance key with the given name.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceKeyDelete.json
 */
async function deleteTheManagedInstanceKey() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const keyName = "someVault_someKey_01234567890123456789012345678901";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstanceKeys.beginDeleteAndWait(
    resourceGroupName,
    managedInstanceName,
    keyName
  );
  console.log(result);
}

deleteTheManagedInstanceKey().catch(console.error);
