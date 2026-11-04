
/**
 * Samples for ReplicationProtectedItems ListByReplicationProtectionContainers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_ListByReplicationProtectionContainers.json
     */
    /**
     * Sample code: Gets the list of Replication protected items.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationProtectedItems(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().listByReplicationProtectionContainers("resourceGroupPS1", "vault1",
            "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", com.azure.core.util.Context.NONE);
    }
}
