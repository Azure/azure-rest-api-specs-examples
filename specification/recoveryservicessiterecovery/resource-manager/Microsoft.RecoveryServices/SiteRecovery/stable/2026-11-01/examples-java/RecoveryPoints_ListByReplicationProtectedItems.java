
/**
 * Samples for RecoveryPoints ListByReplicationProtectedItems.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/RecoveryPoints_ListByReplicationProtectedItems.json
     */
    /**
     * Sample code: Gets the list of recovery points for a replication protected item.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfRecoveryPointsForAReplicationProtectedItem(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.recoveryPoints().listByReplicationProtectedItems("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            com.azure.core.util.Context.NONE);
    }
}
