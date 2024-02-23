
/**
 * Samples for ReplicationProtectionContainers ListByReplicationFabrics.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationProtectionContainers_ListByReplicationFabrics.json
     */
    /**
     * Sample code: Gets the list of protection container for a fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfProtectionContainerForAFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionContainers().listByReplicationFabrics("vault1", "resourceGroupPS1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
