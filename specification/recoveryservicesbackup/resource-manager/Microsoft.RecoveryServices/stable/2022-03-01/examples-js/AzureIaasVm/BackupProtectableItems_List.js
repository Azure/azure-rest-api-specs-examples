const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides a pageable list of protectable objects within your subscription according to the query filter and the
pagination parameters.
 *
 * @summary Provides a pageable list of protectable objects within your subscription according to the query filter and the
pagination parameters.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/BackupProtectableItems_List.json
 */
async function listProtectableItemsWithBackupManagementTypeFilterAsAzureIaasVM() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const filter = "backupManagementType eq 'AzureIaasVM'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupProtectableItems.list(
    vaultName,
    resourceGroupName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProtectableItemsWithBackupManagementTypeFilterAsAzureIaasVM().catch(console.error);
