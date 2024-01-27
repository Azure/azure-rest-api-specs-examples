const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restores the specified backed up data. This is an asynchronous operation. To know the status of this API call, use
GetProtectedItemOperationResult API.
 *
 * @summary Restores the specified backed up data. This is an asynchronous operation. To know the status of this API call, use
GetProtectedItemOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureIaasVm/TriggerRestore_RestoreDisks_IaasVMRestoreRequest.json
 */
async function restoreDisksWithIaasVMRestoreRequest() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "netsdktestrg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const protectedItemName = "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const recoveryPointId = "348916168024334";
  const parameters = {
    properties: {
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
      targetDiskNetworkAccessSettings: {
        targetDiskAccessId:
          "/subscriptions/e7a191f5-713c-4bdb-b5e4-cf3dd90230ef/resourceGroups/arpja/providers/Microsoft.Compute/diskAccesses/arpja-diskaccess-ccy",
        targetDiskNetworkAccessOption: "EnablePrivateAccessForAllDisks",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.restores.beginTriggerAndWait(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName,
    recoveryPointId,
    parameters,
  );
  console.log(result);
}
