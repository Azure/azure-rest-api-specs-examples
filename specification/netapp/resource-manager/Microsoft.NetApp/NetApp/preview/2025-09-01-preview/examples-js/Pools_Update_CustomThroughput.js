const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to patch the specified capacity pool
 *
 * @summary patch the specified capacity pool
 * x-ms-original-file: 2025-09-01-preview/Pools_Update_CustomThroughput.json
 */
async function poolsUpdateCustomThroughput() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.pools.update("myRG", "account1", "customPool1", {});
  console.log(result);
}
