const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function restoreToNewAzureIaasVMWithIaasVMRestoreRequestWithIdentityBasedRestoreDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = "netsdktestrg";
  const fabricName = "Azure";
  const containerName = "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const protectedItemName = "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
  const recoveryPointId = "348916168024334";
  const parameters = {
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
    parameters
  );
  console.log(result);
}

restoreToNewAzureIaasVMWithIaasVMRestoreRequestWithIdentityBasedRestoreDetails().catch(
  console.error
);
