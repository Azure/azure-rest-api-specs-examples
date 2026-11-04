
/**
 * Samples for ReplicationProtectionContainers ListByReplicationFabrics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionContainers_ListByReplicationFabrics.json
     */
    /**
     * Sample code: Gets the list of protection container for a fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfProtectionContainerForAFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainers().listByReplicationFabrics("resourceGroupPS1", "vault1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
