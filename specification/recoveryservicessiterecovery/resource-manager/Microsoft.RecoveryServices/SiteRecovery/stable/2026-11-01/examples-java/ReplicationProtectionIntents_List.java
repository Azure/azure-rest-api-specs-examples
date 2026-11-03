
/**
 * Samples for ReplicationProtectionIntents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionIntents_List.json
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
