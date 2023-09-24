/** Samples for ReplicationJobs Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationJobs_Cancel.json
     */
    /**
     * Sample code: Cancels the specified job.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void cancelsTheSpecifiedJob(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationJobs()
            .cancel(
                "vault1", "resourceGroupPS1", "2653c648-fc72-4316-86f3-fdf8eaa0066b", com.azure.core.util.Context.NONE);
    }
}
