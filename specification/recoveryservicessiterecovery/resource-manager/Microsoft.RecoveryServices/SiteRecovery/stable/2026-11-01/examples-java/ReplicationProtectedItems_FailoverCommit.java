
/**
 * Samples for ReplicationProtectedItems FailoverCommit.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_FailoverCommit.json
     */
    /**
     * Sample code: Execute commit failover.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        executeCommitFailover(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().failoverCommit("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            com.azure.core.util.Context.NONE);
    }
}
