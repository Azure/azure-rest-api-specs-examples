/** Samples for ClusterVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsList_example.json
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
