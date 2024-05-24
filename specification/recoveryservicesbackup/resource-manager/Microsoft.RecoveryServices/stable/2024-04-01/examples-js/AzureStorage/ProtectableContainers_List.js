const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the containers that can be registered to Recovery Services Vault.
 *
 * @summary Lists the containers that can be registered to Recovery Services Vault.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/AzureStorage/ProtectableContainers_List.json
 */
async function listProtectableItemsWithBackupManagementTypeFilterAsAzureStorage() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "testvault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "testRg";
  const fabricName = "Azure";
  const filter = "backupManagementType eq 'AzureStorage' and workloadType eq 'AzureFileShare'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.protectableContainers.list(
    vaultName,
    resourceGroupName,
    fabricName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
