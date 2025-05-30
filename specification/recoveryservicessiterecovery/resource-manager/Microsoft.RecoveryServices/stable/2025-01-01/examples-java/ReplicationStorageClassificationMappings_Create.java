
import com.azure.resourcemanager.recoveryservicessiterecovery.models.StorageMappingInputProperties;

/**
 * Samples for StorageClassificationMappings Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationStorageClassificationMappings_Create.json
     */
    /**
     * Sample code: Create storage classification mapping.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void createStorageClassificationMapping(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.storageClassificationMappings().define("testStorageMapping")
            .withExistingReplicationStorageClassification("resourceGroupPS1", "vault1",
                "2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0",
                "8891569e-aaef-4a46-a4a0-78c14f2d7b09")
            .withProperties(new StorageMappingInputProperties().withTargetStorageClassificationId(
                "/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0/replicationStorageClassifications/8891569e-aaef-4a46-a4a0-78c14f2d7b09"))
            .create();
    }
}
