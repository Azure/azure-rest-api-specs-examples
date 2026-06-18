
/**
 * Samples for Services Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ServiceGetOperation_example.json
     */
    /**
     * Sample code: Get a service.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getAService(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.services().getWithResponse("resRg", "myCluster", "myApp", "myService",
            com.azure.core.util.Context.NONE);
    }
}
