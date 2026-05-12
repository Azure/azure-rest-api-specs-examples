const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to restores the specified backed up data. This is an asynchronous operation. To know the status of this API call, use
 * GetProtectedItemOperationResult API.
 *
 * @summary restores the specified backed up data. This is an asynchronous operation. To know the status of this API call, use
 * GetProtectedItemOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/TriggerRestore_RestoreDisks_IaasVMRestoreWithRehydrationRequest.json
 */
async function restoreDisksWithIaasVMRestoreWithRehydrationRequest() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  await client.restores.trigger(
    "testVault",
    "netsdktestrg",
    "Azure",
    "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
    "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
    "348916168024334",
    {
      properties: {
        createNewCloudService: true,
        encryptionDetails: { encryptionEnabled: false },
        objectType: "IaasVMRestoreWithRehydrationRequest",
        originalStorageAccountOption: false,
        recoveryPointId: "348916168024334",
        recoveryPointRehydrationInfo: {
          rehydrationPriority: "Standard",
          rehydrationRetentionDuration: "P7D",
        },
        recoveryType: "RestoreDisks",
        region: "southeastasia",
        sourceResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1",
        storageAccountId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount",
      },
    },
  );
}
