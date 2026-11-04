
/**
 * Samples for ReplicationFabrics MigrateToAad.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationFabrics_MigrateToAad.json
     */
    /**
     * Sample code: Migrates the site to AAD.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        migratesTheSiteToAAD(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().migrateToAad("resourceGroupPS1", "vault1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
