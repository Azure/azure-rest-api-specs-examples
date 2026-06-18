
/**
 * Samples for ClusterVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ClusterVersionsList_example.json
     */
    /**
     * Sample code: List cluster versions.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void listClusterVersions(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusterVersions().listWithResponse("eastus", com.azure.core.util.Context.NONE);
    }
}
