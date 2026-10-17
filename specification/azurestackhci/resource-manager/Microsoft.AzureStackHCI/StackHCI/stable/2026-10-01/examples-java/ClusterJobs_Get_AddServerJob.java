
/**
 * Samples for ClusterJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ClusterJobs_Get_AddServerJob.json
     */
    /**
     * Sample code: ClusterJobs_Get_AddServerJob.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        clusterJobsGetAddServerJob(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusterJobs().getWithResponse("test-rg", "myCluster", "addserver", com.azure.core.util.Context.NONE);
    }
}
