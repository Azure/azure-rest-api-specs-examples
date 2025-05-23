
/**
 * Samples for ReplicationProtectionClusters Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionClusters_Get.json
     */
    /**
     * Sample code: Gets the details of a Replication protection cluster.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDetailsOfAReplicationProtectionCluster(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionClusters().getWithResponse("resourceGroupPS1", "vault1", "eastus",
            "eastus-container", "cluster1", com.azure.core.util.Context.NONE);
    }
}
