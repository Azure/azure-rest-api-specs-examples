const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets a backup instance with name in a backup vault
 *
 * @summary Gets a backup instance with name in a backup vault
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/BackupInstanceOperations/GetBackupInstance_ADLSBlobBackupDatasourceParameters.json
 */
async function getBackupInstanceForAdlsBlob() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "54707983-993e-43de-8d94-074451394eda";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "adlsrg";
  const vaultName = "adlsvault";
  const backupInstanceName = "adlsbackupinstance";
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.get(resourceGroupName, vaultName, backupInstanceName);
  console.log(result);
}
