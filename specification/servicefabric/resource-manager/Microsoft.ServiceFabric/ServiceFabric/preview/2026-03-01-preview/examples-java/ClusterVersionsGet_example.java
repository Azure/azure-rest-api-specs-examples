
/**
 * Samples for ClusterVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ClusterVersionsGet_example.json
     */
    /**
     * Sample code: Get cluster version.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getClusterVersion(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusterVersions().getWithResponse("eastus", "6.1.480.9494", com.azure.core.util.Context.NONE);
    }
}
