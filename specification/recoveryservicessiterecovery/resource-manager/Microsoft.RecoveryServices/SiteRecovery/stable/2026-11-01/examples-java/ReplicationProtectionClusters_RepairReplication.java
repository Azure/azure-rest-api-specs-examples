
/**
 * Samples for ReplicationProtectionClusters RepairReplication.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionClusters_RepairReplication.json
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
