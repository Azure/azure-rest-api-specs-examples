
/**
 * Samples for ApplicationTypes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ApplicationTypeNameGetOperation_example.json
     */
    /**
     * Sample code: Get an application type.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAnApplicationType(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypes().getWithResponse("resRg", "myCluster", "myAppType", com.azure.core.util.Context.NONE);
    }
}
