const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validate operation for specified backed up item. This is a synchronous operation.
 *
 * @summary Validate operation for specified backed up item. This is a synchronous operation.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/ValidateOperation_RestoreDisk_IdentityBasedRestoreDetails.json
 */
async function validateOperationWithIdentityBasedRestoreDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "testRG";
  const parameters = {
    objectType: "ValidateIaasVMRestoreOperationRequest",
    restoreRequest: {
      createNewCloudService: true,
      encryptionDetails: { encryptionEnabled: false },
      identityBasedRestoreDetails: {
        targetStorageAccountId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount",
      },
      identityInfo: {
        isSystemAssignedIdentity: false,
        managedIdentityResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi",
      },
      objectType: "IaasVMRestoreRequest",
      originalStorageAccountOption: false,
      recoveryPointId: "348916168024334",
      recoveryType: "RestoreDisks",
      region: "southeastasia",
      sourceResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.operation.validate(vaultName, resourceGroupName, parameters);
  console.log(result);
}

validateOperationWithIdentityBasedRestoreDetails().catch(console.error);
