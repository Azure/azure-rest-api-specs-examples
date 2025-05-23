
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ClusterTestFailoverCleanupInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ClusterTestFailoverCleanupInputProperties;

/**
 * Samples for ReplicationProtectionClusters TestFailoverCleanup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionClusters_TestFailoverCleanup.json
     */
    /**
     * Sample code: Execute test failover cleanup for cluster.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeTestFailoverCleanupForCluster(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionClusters().testFailoverCleanup("resourceGroupPS1", "vault1", "fabric-pri-eastus",
            "pri-cloud-eastus", "testcluster",
            new ClusterTestFailoverCleanupInput()
                .withProperties(new ClusterTestFailoverCleanupInputProperties().withComments("Test Failover Cleanup")),
            com.azure.core.util.Context.NONE);
    }
}
