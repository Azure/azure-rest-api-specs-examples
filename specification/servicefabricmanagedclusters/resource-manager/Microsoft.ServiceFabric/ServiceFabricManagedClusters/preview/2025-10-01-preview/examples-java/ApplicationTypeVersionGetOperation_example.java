
/**
 * Samples for ApplicationTypeVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ApplicationTypeVersionGetOperation_example.json
     */
    /**
     * Sample code: Get an application type version.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getAnApplicationTypeVersion(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applicationTypeVersions().getWithResponse("resRg", "myCluster", "myAppType", "1.0",
            com.azure.core.util.Context.NONE);
    }
}
