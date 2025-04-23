
/**
 * Samples for ClusterRecoveryPoints ListByReplicationProtectionCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ClusterRecoveryPoints_ListByReplicationProtectionCluster.json
     */
    /**
     * Sample code: Gets the list of cluster recovery points.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfClusterRecoveryPoints(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.clusterRecoveryPoints().listByReplicationProtectionCluster("resourceGroupPS1", "vault1",
            "fabric-pri-eastus", "pri-cloud-eastus", "testcluster", com.azure.core.util.Context.NONE);
    }
}
