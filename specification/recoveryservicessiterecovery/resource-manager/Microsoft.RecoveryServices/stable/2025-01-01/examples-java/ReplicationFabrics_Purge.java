
/**
 * Samples for ReplicationFabrics Purge.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationFabrics_Purge.json
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
