
/**
 * Samples for ReplicationProtectionClusters RepairReplication.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectionClusters_RepairReplication.json
     */
    /**
     * Sample code: Resynchronize or repair replication of protection cluster.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void resynchronizeOrRepairReplicationOfProtectionCluster(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionClusters().repairReplication("resourceGroupPS1", "vault1", "eastus",
            "eastus-container", "cluster12", com.azure.core.util.Context.NONE);
    }
}
