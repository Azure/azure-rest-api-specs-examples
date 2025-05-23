
/**
 * Samples for StorageClassifications ListByReplicationFabrics.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationStorageClassifications_ListByReplicationFabrics.json
     */
    /**
     * Sample code: Gets the list of storage classification objects under a fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfStorageClassificationObjectsUnderAFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.storageClassifications().listByReplicationFabrics("resourceGroupPS1", "vault1",
            "2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0", com.azure.core.util.Context.NONE);
    }
}
