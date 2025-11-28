
/**
 * Samples for Services Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ServiceGetOperation_example.json
     */
    /**
     * Sample code: Get a service.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAService(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.services().getWithResponse("resRg", "myCluster", "myApp", "myService",
            com.azure.core.util.Context.NONE);
    }
}
