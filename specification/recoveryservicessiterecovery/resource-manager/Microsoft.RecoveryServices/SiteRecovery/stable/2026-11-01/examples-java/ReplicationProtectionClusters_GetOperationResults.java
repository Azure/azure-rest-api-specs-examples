
/**
 * Samples for ReplicationProtectionClusters GetOperationResults.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectionClusters_GetOperationResults.json
     */
    /**
     * Sample code: Tracks the Replication protection cluster async operation.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void tracksTheReplicationProtectionClusterAsyncOperation(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectionClusters().getOperationResultsWithResponse("resourceGroupPS1", "vault1", "eastus",
            "eastus-container", "cluster1", "ea63a935-59d5-4b12-aff2-98773f63c116", com.azure.core.util.Context.NONE);
    }
}
