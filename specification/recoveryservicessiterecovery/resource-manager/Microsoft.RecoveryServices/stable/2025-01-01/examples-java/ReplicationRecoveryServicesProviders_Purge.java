
/**
 * Samples for ReplicationRecoveryServicesProviders Purge.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationRecoveryServicesProviders_Purge.json
     */
    /**
     * Sample code: Purges recovery service provider from fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void purgesRecoveryServiceProviderFromFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryServicesProviders().purge("resourceGroupPS1", "vault1", "cloud1",
            "241641e6-ee7b-4ee4-8141-821fadda43fa", com.azure.core.util.Context.NONE);
    }
}
