
/**
 * Samples for Clusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/GetCluster.json
     */
    /**
     * Sample code: Get cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getCluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().getByResourceGroupWithResponse("test-rg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
