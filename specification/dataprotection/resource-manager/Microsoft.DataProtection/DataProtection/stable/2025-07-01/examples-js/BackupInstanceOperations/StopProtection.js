const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation will stop protection of a backup instance and data will be held forever
 *
 * @summary this operation will stop protection of a backup instance and data will be held forever
 * x-ms-original-file: 2025-07-01/BackupInstanceOperations/StopProtection.json
 */
async function stopProtection() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const client = new DataProtectionClient(credential, subscriptionId);
  await client.backupInstances.stopProtection("testrg", "testvault", "testbi");
}
