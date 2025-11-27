
/**
 * Samples for ManagedClusterVersion List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ManagedClusterVersionList_example.json
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
