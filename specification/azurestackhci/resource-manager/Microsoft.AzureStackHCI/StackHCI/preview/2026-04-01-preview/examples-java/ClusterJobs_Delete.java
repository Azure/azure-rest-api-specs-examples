
/**
 * Samples for ClusterJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ClusterJobs_Delete.json
     */
    /**
     * Sample code: ClusterJobs_Delete.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void clusterJobsDelete(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusterJobs().delete("rghci", "3-Vz3LSRO5Q6q8EV-PKs8-5E", "configureSdnIntegration",
            com.azure.core.util.Context.NONE);
    }
}
