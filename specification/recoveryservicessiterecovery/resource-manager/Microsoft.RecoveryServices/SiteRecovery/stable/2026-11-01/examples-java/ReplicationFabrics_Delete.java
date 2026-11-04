
/**
 * Samples for ReplicationFabrics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationFabrics_Delete.json
     */
    /**
     * Sample code: Deletes the site.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        deletesTheSite(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().delete("resourceGroupPS1", "vault1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
