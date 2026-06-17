
/**
 * Samples for Applications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ApplicationGetOperation_example.json
     */
    /**
     * Sample code: Get an application.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getAnApplication(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applications().getWithResponse("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
