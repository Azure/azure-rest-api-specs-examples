const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides a pageable list of all items that are backed up within a vault.
 *
 * @summary Provides a pageable list of all items that are backed up within a vault.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/BackupProtectedItems_List.json
 */
async function listProtectedItemsWithBackupManagementTypeFilterAsAzureIaasVM() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const filter = "backupManagementType eq 'AzureIaasVM' and itemType eq 'VM'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupProtectedItems.list(vaultName, resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProtectedItemsWithBackupManagementTypeFilterAsAzureIaasVM().catch(console.error);
