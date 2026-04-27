
/**
 * Samples for Clusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/DeleteCluster.json
     */
    /**
     * Sample code: Delete cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteCluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().delete("test-rg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
