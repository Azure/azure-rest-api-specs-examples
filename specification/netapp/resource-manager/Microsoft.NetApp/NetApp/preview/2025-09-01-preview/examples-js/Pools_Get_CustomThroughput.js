const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get details of the specified capacity pool
 *
 * @summary get details of the specified capacity pool
 * x-ms-original-file: 2025-09-01-preview/Pools_Get_CustomThroughput.json
 */
async function poolsGetCustomThroughput() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.pools.get("myRG", "account1", "customPool1");
  console.log(result);
}
