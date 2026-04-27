
/**
 * Samples for Clusters CreateIdentity.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CreateClusterIdentity.json
     */
    /**
     * Sample code: Create cluster Identity.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createClusterIdentity(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().createIdentity("test-rg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
