const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets information about the specified pool.
 *
 * @summary gets information about the specified pool.
 * x-ms-original-file: 2025-06-01/PoolGet_VirtualMachineConfiguration_MangedOSDisk.json
 */
async function getPoolVirtualMachineConfigurationOSDisk() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.get("default-azurebatch-japaneast", "sampleacct", "testpool");
  console.log(result);
}
