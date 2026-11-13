const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to update the recovery settings of an ASR replication protected item.
 *
 * @summary the operation to update the recovery settings of an ASR replication protected item.
 * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Update_A2ACvm.json
 */
async function updatesTheTargetConfidentialDiskEncryptionSetForAnA2AConfidentialVM() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "42195872-7e70-4f8a-837f-84b28ecbb78b";
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectedItems.update(
    "priyanprg",
    "priyanponeboxvault",
    "CentralUSCanSite",
    "CentralUSCancloud",
    "a2aCvmVm",
    {
      properties: {
        providerSpecificDetails: {
          instanceType: "A2A",
          recoveryResourceGroupId:
            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg",
          managedDiskUpdateDetails: [
            {
              diskId:
                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/disks/a2aCvmVm-osdisk",
              recoveryReplicaDiskAccountType: "Premium_LRS",
              recoveryTargetDiskAccountType: "Premium_LRS",
              confidentialDiskEncryptionInfo: {
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
