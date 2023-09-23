/** Samples for ReplicationProtectedItems FailoverCommit. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationProtectedItems_FailoverCommit.json
     */
    /**
     * Sample code: Execute commit failover.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeCommitFailover(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationProtectedItems()
            .failoverCommit(
                "vault1",
                "resourceGroupPS1",
                "cloud1",
                "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
                "f8491e4f-817a-40dd-a90c-af773978c75b",
                com.azure.core.util.Context.NONE);
    }
}
