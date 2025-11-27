
/**
 * Samples for Services ListByApplications.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ServiceListOperation_example.json
     */
    /**
     * Sample code: Get a list of service resources.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAListOfServiceResources(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.services().listByApplications("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
