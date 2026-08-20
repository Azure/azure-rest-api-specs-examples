const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to restores the specified backed up data. This is an asynchronous operation. To know the status of this API call, use
 * GetProtectedItemOperationResult API.
 *
 * @summary restores the specified backed up data. This is an asynchronous operation. To know the status of this API call, use
 * GetProtectedItemOperationResult API.
 * x-ms-original-file: 2026-07-01/AzureIaasVm/TriggerRestore_ALR_IaasVMRestoreRequest_DataDiskEncryption_SingleEntry.json
 */
async function restoreToNewAzureIaasVmWithIaasVMRestoreRequestWithIdentityBasedRestoreDetailsAndDataDiskEncryptionSettings() {
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
        createNewCloudService: false,
        encryptionDetails: { encryptionEnabled: false },
        identityBasedRestoreDetails: {
          targetStorageAccountId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Storage/storageAccounts/testingAccount",
        },
        identityInfo: { isSystemAssignedIdentity: true },
        objectType: "IaasVMRestoreRequest",
        originalStorageAccountOption: false,
        recoveryPointId: "348916168024334",
        recoveryType: "AlternateLocation",
        region: "southeastasia",
        securedVMDetails: {
          securedVMOsDiskEncryptionSetId:
            "/subscriptions/5288acd1-ba79-4377-9205-9f220331a44a/resourceGroups/asmaskarrg-cvm-os-cmk-3cmk-1none-data-2606040706/providers/Microsoft.Compute/diskEncryptionSets/des-os",
          dataDiskEncryptionSettings: {
            dataDiskEncryptionSetId:
              "/subscriptions/5288acd1-ba79-4377-9205-9f220331a44a/resourceGroups/asmaskarrg-cvm-os-cmk-3cmk-1none-data-2606040706/providers/Microsoft.Compute/diskEncryptionSets/des-data-lun-0",
            dataDiskEncryptionIdentity:
              "/subscriptions/5288acd1-ba79-4377-9205-9f220331a44a/resourcegroups/asmaskarrg-cvm-os-cmk-3cmk-1none-data-2606040706/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami-cdde",
          },
        },
        sourceResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1",
        subnetId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet/subnets/default",
        targetResourceGroupId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2",
        targetVirtualMachineId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2/providers/Microsoft.Compute/virtualmachines/RSMDALRVM981435",
        virtualNetworkId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet",
      },
    },
  );
}
