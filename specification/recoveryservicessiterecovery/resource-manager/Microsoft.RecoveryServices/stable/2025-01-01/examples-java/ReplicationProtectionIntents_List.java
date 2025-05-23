
/**
 * Samples for ReplicationProtectionIntents List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionIntents_List.json
     */
    /**
     * Sample code: Gets the list of replication protection intent objects.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationProtectionIntentObjects(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionIntents().list("resourceGroupPS1", "2007vttp", null, null,
            com.azure.core.util.Context.NONE);
    }
}
