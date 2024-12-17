
/**
 * Samples for ManagedClusterVersion Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/ManagedClusterVersionGet_example.json
     */
    /**
     * Sample code: Get cluster version.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getClusterVersion(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusterVersions().getWithResponse("eastus", "7.2.477.9590", com.azure.core.util.Context.NONE);
    }
}
