
/**
 * Samples for ReplicationProtectedItems List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_List.json
     */
    /**
     * Sample code: Gets the list of replication protected items.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationProtectedItems(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().list("resourceGroupPS1", "vault1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
