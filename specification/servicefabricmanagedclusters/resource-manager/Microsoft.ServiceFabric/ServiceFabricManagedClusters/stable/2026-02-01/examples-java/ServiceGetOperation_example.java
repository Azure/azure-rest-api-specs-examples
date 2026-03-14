
/**
 * Samples for Services Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ServiceGetOperation_example.json
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
