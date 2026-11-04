
/**
 * Samples for ReplicationProtectionClusters FailoverCommit.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionClusters_FailoverCommit.json
     */
    /**
     * Sample code: Execute commit failover for cluster.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeCommitFailoverForCluster(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionClusters().failoverCommit("resourceGroupPS1", "vault1", "fabric-pri-eastus",
            "pri-cloud-eastus", "testcluster", com.azure.core.util.Context.NONE);
    }
}
