
/**
 * Samples for ClusterJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ClusterJobs_Get_ConfigureSdnIntegrationJob.json
     */
    /**
     * Sample code: ClusterJobs_Get_ConfigureSdnIntegrationJob.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        clusterJobsGetConfigureSdnIntegrationJob(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusterJobs().getWithResponse("rghci", "Y-k0MG", "configureSdnIntegration",
            com.azure.core.util.Context.NONE);
    }
}
