
/**
 * Samples for ReplicationJobs Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationJobs_Restart.json
     */
    /**
     * Sample code: Restarts the specified job.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        restartsTheSpecifiedJob(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationJobs().restart("resourceGroupPS1", "vault1", "0664564c-353e-401a-ab0c-722257c10e25",
            com.azure.core.util.Context.NONE);
    }
}
