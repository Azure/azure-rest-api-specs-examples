const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists of backup policies associated with Recovery Services Vault. API provides pagination parameters to fetch
scoped results.
 *
 * @summary Lists of backup policies associated with Recovery Services Vault. API provides pagination parameters to fetch
scoped results.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureWorkload/BackupPolicies_List.json
 */
async function listProtectionPoliciesWithBackupManagementTypeFilterAsAzureWorkload() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const filter = "backupManagementType eq 'AzureWorkload'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupPolicies.list(vaultName, resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProtectionPoliciesWithBackupManagementTypeFilterAsAzureWorkload().catch(console.error);
