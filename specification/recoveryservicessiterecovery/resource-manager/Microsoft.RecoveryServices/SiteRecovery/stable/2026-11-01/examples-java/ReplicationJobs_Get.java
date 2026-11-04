
/**
 * Samples for ReplicationJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationJobs_Get.json
     */
    /**
     * Sample code: Gets the job details.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheJobDetails(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationJobs().getWithResponse("resourceGroupPS1", "vault1", "58776d0b-3141-48b2-a377-9ad863eb160d",
            com.azure.core.util.Context.NONE);
    }
}
