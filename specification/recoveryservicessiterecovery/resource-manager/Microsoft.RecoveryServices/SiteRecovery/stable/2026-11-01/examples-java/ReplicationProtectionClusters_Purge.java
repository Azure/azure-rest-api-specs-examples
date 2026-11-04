
/**
 * Samples for ReplicationProtectionClusters Purge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionClusters_Purge.json
     */
    /**
     * Sample code: Purge the replication protection cluster.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void purgeTheReplicationProtectionCluster(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionClusters().purge("resourceGroupPS1", "vault1", "eastus", "eastus-container",
            "cluster1", com.azure.core.util.Context.NONE);
    }
}
