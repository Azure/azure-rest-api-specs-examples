
/**
 * Samples for ReplicationProtectedItems FailoverCancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationProtectedItems_FailoverCancel.json
     */
    /**
     * Sample code: Execute cancel failover.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        executeCancelFailover(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().failoverCancel("resourceGroupPS1", "vault1", "cloud1",
            "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179", "f8491e4f-817a-40dd-a90c-af773978c75b",
            com.azure.core.util.Context.NONE);
    }
}
