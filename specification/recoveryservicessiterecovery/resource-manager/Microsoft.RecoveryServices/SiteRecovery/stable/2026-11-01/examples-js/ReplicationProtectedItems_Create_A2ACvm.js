const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create an ASR replication protected item (Enable replication).
 *
 * @summary the operation to create an ASR replication protected item (Enable replication).
 * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Create_A2ACvm.json
 */
async function enablesProtectionForAnA2AConfidentialVMWithCMKDiskEncryption() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "42195872-7e70-4f8a-837f-84b28ecbb78b";
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectedItems.create(
    "priyanprg",
    "priyanponeboxvault",
    "CentralUSCanSite",
    "CentralUSCancloud",
    "a2aCvmVm",
    {
      properties: {
        policyId:
          "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationPolicies/a2aPolicy",
        protectableItemId:
          "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/CentralUSCanSite/replicationProtectionContainers/CentralUSCancloud/replicationProtectableItems/a2aCvmVm",
        providerSpecificDetails: {
          instanceType: "A2A",
          fabricObjectId:
            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/virtualMachines/a2aCvmVm",
          recoveryContainerId:
            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/EUSCanSite/replicationProtectionContainers/euscancloud",
          recoveryResourceGroupId:
            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg",
          vmManagedDisks: [
            {
              diskId:
                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/disks/a2aCvmVm-osdisk",
              primaryStagingAzureStorageAccountId:
                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Storage/storageAccounts/a2acache",
              recoveryResourceGroupId:
                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg",
              recoveryReplicaDiskAccountType: "Premium_LRS",
              recoveryTargetDiskAccountType: "Premium_LRS",
              confidentialDiskEncryptionInfo: {
                recoveryReplicaConfidentialDiskEncryptionSetId:
                  "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-replica-des",
                recoveryTargetConfidentialDiskEncryptionSetId:
                  "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-target-des",
              },
            },
          ],
          recoveryConfidentialDataDiskEncryptionIdentity:
            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/cvm-cdde-uami",
        },
      },
    },
  );
  console.log(result);
}
