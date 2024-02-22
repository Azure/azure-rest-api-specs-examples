
/**
 * Samples for ReplicationProtectedItems List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationProtectedItems_List.json
     */
    /**
     * Sample code: Gets the list of replication protected items.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationProtectedItems(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().list("vault1", "resourceGroupPS1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
