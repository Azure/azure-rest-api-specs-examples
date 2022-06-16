const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function validateOperation() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "testRG";
  const parameters = {
    objectType: "ValidateIaasVMRestoreOperationRequest",
    restoreRequest: {
      createNewCloudService: true,
      encryptionDetails: { encryptionEnabled: false },
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
      storageAccountId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.operation.validate(vaultName, resourceGroupName, parameters);
  console.log(result);
}

validateOperation().catch(console.error);
