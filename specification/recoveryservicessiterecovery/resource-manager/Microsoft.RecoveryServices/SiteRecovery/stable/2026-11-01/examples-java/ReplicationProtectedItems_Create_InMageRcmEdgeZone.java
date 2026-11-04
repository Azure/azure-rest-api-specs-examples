
import com.azure.resourcemanager.recoveryservicessiterecovery.models.EnableProtectionInputProperties;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ExtendedLocation;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ExtendedLocationType;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.InMageRcmEnableProtectionInput;

/**
 * Samples for ReplicationProtectedItems Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Create_InMageRcmEdgeZone.json
     */
    /**
     * Sample code: Enables protection for an InMage RCM VM in an extended location.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void enablesProtectionForAnInMageRCMVMInAnExtendedLocation(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().define("machine1")
            .withExistingReplicationProtectionContainer("resourceGroupPS1", "vault1", "vmware-fabric",
                "vmware-container")
            .withProperties(new EnableProtectionInputProperties().withPolicyId(
                "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/protectionProfile1")
                .withProtectableItemId(
                    "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/vmware-fabric/replicationProtectionContainers/vmware-container/replicationProtectableItems/machine1")
                .withProviderSpecificDetails(new InMageRcmEnableProtectionInput().withFabricDiscoveryMachineId(
                    "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/sourceResourceGroup/providers/Microsoft.OffAzure/VMwareSites/site1/machines/machine1")
                    .withTargetResourceGroupId(
                        "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/targetResourceGroup")
                    .withTargetNetworkId(
                        "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/targetResourceGroup/providers/Microsoft.Network/virtualNetworks/targetVNet")
                    .withTargetSubnetName("default").withTargetVmName("machine1").withTargetVmSize("Standard_D2s_v3")
                    .withProcessServerId("process-server-1").withTargetExtendedLocation(
                        new ExtendedLocation().withName("edge-zone-1").withType(ExtendedLocationType.EDGE_ZONE))))
            .create();
    }
}
