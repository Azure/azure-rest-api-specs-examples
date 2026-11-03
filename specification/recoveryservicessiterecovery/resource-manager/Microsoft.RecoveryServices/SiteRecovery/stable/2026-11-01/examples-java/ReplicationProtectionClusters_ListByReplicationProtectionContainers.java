
/**
 * Samples for ReplicationProtectionClusters ListByReplicationProtectionContainers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionClusters_ListByReplicationProtectionContainers.json
     */
    /**
     * Sample code: Gets the list of Replication protection clusters in fabric, container.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfReplicationProtectionClustersInFabricContainer(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionClusters().listByReplicationProtectionContainers("resourceGroupPS1", "vault1",
            "eastus", "eastus-container", com.azure.core.util.Context.NONE);
    }
}
