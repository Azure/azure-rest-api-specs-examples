
/**
 * Samples for ReplicationLogicalNetworks ListByReplicationFabrics.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationLogicalNetworks_ListByReplicationFabrics.json
     */
    /**
     * Sample code: Gets the list of logical networks under a fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfLogicalNetworksUnderAFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationLogicalNetworks().listByReplicationFabrics("vault1", "resourceGroupPS1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
