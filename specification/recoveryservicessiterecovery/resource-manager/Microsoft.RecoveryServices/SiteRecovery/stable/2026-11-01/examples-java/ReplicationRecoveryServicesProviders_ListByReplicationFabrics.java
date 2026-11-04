
/**
 * Samples for ReplicationRecoveryServicesProviders ListByReplicationFabrics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationRecoveryServicesProviders_ListByReplicationFabrics.json
     */
    /**
     * Sample code: Gets the list of registered recovery services providers for the fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfRegisteredRecoveryServicesProvidersForTheFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryServicesProviders().listByReplicationFabrics("resourceGroupPS1", "vault1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
