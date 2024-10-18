
/**
 * Samples for ManagedClusterVersion List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/
     * examples/ManagedClusterVersionList_example.json
     */
    /**
     * Sample code: List cluster versions.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listClusterVersions(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusterVersions().listWithResponse("eastus", com.azure.core.util.Context.NONE);
    }
}
