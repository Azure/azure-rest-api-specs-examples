const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create an ASR replication protected item (Enable replication).
 *
 * @summary the operation to create an ASR replication protected item (Enable replication).
 * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Create_InMageRcmEdgeZone.json
 */
async function enablesProtectionForAnInMageRCMVMInAnExtendedLocation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectedItems.create(
    "resourceGroupPS1",
    "vault1",
    "vmware-fabric",
    "vmware-container",
    "machine1",
    {
      properties: {
        policyId:
          "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionProfile1",
        protectableItemId:
          "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/vmware-fabric/replicationProtectionContainers/vmware-container/replicationProtectableItems/machine1",
        providerSpecificDetails: {
          fabricDiscoveryMachineId:
            "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/sourceResourceGroup/providers/Microsoft.OffAzure/VMwareSites/site1/machines/machine1",
          instanceType: "InMageRcm",
          processServerId: "process-server-1",
          targetExtendedLocation: { name: "edge-zone-1", type: "EdgeZone" },
          targetNetworkId:
            "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/targetResourceGroup/providers/Microsoft.Network/virtualNetworks/targetVNet",
          targetResourceGroupId:
            "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/targetResourceGroup",
          targetSubnetName: "default",
          targetVmName: "machine1",
          targetVmSize: "Standard_D2s_v3",
        },
      },
    },
  );
  console.log(result);
}
