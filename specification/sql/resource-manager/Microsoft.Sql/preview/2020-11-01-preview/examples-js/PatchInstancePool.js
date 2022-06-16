const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an instance pool.
 *
 * @summary Updates an instance pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/PatchInstancePool.json
 */
async function patchAnInstancePool() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const instancePoolName = "testIP";
  const parameters = { tags: { x: "y" } };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.instancePools.beginUpdateAndWait(
    resourceGroupName,
    instancePoolName,
    parameters
  );
  console.log(result);
}

patchAnInstancePool().catch(console.error);
