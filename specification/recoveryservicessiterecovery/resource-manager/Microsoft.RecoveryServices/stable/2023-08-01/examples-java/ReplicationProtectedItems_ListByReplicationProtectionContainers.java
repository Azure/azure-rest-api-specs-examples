
/**
 * Samples for ReplicationProtectedItems ListByReplicationProtectionContainers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationProtectedItems_ListByReplicationProtectionContainers.json
     */
    /**
     * Sample code: Gets the list of Replication protected items.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationProtectedItems(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().listByReplicationProtectionContainers("vault1", "resourceGroupPS1",
            "cloud1", "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", com.azure.core.util.Context.NONE);
    }
}
