/** Samples for ReplicationFabrics MigrateToAad. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationFabrics_MigrateToAad.json
     */
    /**
     * Sample code: Migrates the site to AAD.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void migratesTheSiteToAAD(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationFabrics()
            .migrateToAad("vault1", "resourceGroupPS1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
