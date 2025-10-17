const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a backup instance with name in a backup vault
 *
 * @summary gets a backup instance with name in a backup vault
 * x-ms-original-file: 2025-07-01/BackupInstanceOperations/GetBackupInstance_ADLSBlobBackupDatasourceParameters.json
 */
async function getBackupInstanceForAdlsBlob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "54707983-993e-43de-8d94-074451394eda";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.get("adlsrg", "adlsvault", "adlsbackupinstance");
  console.log(result);
}
