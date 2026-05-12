const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists of backup policies associated with Recovery Services Vault. API provides pagination parameters to fetch
 * scoped results.
 *
 * @summary lists of backup policies associated with Recovery Services Vault. API provides pagination parameters to fetch
 * scoped results.
 * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/BackupPolicies_List.json
 */
async function listProtectionPoliciesWithBackupManagementTypeFilterAsAzureIaasVm() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.backupPolicies.list("NetSDKTestRsVault", "SwaggerTestRg", {
    filter: "backupManagementType eq 'AzureIaasVM'",
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
