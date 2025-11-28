
/**
 * Samples for ApplicationTypeVersions ListByApplicationTypes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ApplicationTypeVersionListOperation_example.json
     */
    /**
     * Sample code: Get a list of application type version resources.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAListOfApplicationTypeVersionResources(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypeVersions().listByApplicationTypes("resRg", "myCluster", "myAppType",
            com.azure.core.util.Context.NONE);
    }
}
