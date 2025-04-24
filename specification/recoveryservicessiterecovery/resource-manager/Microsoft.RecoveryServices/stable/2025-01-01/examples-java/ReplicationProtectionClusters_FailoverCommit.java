
/**
 * Samples for ReplicationProtectionClusters FailoverCommit.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionClusters_FailoverCommit.json
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
