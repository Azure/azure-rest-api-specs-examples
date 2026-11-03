
/**
 * Samples for ReplicationFabrics Purge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationFabrics_Purge.json
     */
    /**
     * Sample code: Purges the site.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        purgesTheSite(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().purge("resourceGroupPS1", "vault1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
