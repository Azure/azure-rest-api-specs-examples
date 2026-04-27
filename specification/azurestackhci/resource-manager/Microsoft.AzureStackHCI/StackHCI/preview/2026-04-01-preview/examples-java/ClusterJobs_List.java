
/**
 * Samples for ClusterJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ClusterJobs_List.json
     */
    /**
     * Sample code: ClusterJobs_List.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void clusterJobsList(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusterJobs().list("rghci", "Ql40O4-I77S", com.azure.core.util.Context.NONE);
    }
}
