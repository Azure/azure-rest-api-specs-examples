
/**
 * Samples for ReplicationJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationJobs_Get.json
     */
    /**
     * Sample code: Gets the job details.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheJobDetails(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationJobs().getWithResponse("vault1", "resourceGroupPS1", "58776d0b-3141-48b2-a377-9ad863eb160d",
            com.azure.core.util.Context.NONE);
    }
}
