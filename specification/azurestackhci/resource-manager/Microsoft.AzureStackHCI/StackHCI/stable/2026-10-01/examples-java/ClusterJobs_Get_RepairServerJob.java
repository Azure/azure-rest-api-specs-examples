
/**
 * Samples for ClusterJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ClusterJobs_Get_RepairServerJob.json
     */
    /**
     * Sample code: ClusterJobs_Get_RepairServerJob.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        clusterJobsGetRepairServerJob(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusterJobs().getWithResponse("test-rg", "myCluster", "repairserver", com.azure.core.util.Context.NONE);
    }
}
