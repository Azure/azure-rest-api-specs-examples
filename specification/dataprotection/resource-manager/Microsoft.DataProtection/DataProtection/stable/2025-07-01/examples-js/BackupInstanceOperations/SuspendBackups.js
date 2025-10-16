const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to this operation will stop backup for a backup instance and retains the backup data as per the policy (except latest Recovery point, which will be retained forever)
 *
 * @summary this operation will stop backup for a backup instance and retains the backup data as per the policy (except latest Recovery point, which will be retained forever)
 * x-ms-original-file: 2025-07-01/BackupInstanceOperations/SuspendBackups.json
 */
async function suspendBackups() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const client = new DataProtectionClient(credential, subscriptionId);
  await client.backupInstances.suspendBackups("testrg", "testvault", "testbi");
}
