const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Used to remove intent from an item
 *
 * @summary Used to remove intent from an item
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureWorkload/BackupProtectionIntent_Delete.json
 */
async function deleteProtectionIntentFromItem() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "myVault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "myRG";
  const fabricName = "Azure";
  const intentObjectName = "249D9B07-D2EF-4202-AA64-65F35418564E";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionIntentOperations.delete(
    vaultName,
    resourceGroupName,
    fabricName,
    intentObjectName,
  );
  console.log(result);
}
