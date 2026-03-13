
/**
 * Samples for Applications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ApplicationGetOperation_example.json
     */
    /**
     * Sample code: Get an application.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAnApplication(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applications().getWithResponse("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
