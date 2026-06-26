const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a managed instance.
 *
 * @summary updates a managed instance.
 * x-ms-original-file: 2025-01-01/ManagedInstanceUpdateMin.json
 */
async function updateManagedInstanceWithMinimalProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstances.update("testrg", "testinstance", {
    tags: { tagKey1: "TagValue1" },
  });
  console.log(result);
}
