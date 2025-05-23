
/**
 * Samples for ReplicationProtectedItems Purge.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationProtectedItems_Purge.json
     */
    /**
     * Sample code: Purges protection.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        purgesProtection(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().purge("vault1", "resourceGroupPS1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "c0c14913-3d7a-48ea-9531-cc99e0e686e6",
            com.azure.core.util.Context.NONE);
    }
}
