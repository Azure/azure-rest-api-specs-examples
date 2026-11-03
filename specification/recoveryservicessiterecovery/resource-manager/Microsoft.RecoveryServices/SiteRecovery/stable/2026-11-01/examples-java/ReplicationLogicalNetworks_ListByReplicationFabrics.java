
/**
 * Samples for ReplicationLogicalNetworks ListByReplicationFabrics.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationLogicalNetworks_ListByReplicationFabrics.json
     */
    /**
     * Sample code: Gets the list of logical networks under a fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfLogicalNetworksUnderAFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationLogicalNetworks().listByReplicationFabrics("resourceGroupPS1", "vault1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
