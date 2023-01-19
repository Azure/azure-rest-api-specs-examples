import com.azure.resourcemanager.servicefabric.models.ClusterVersionsEnvironment;

/** Samples for ClusterVersions ListByEnvironment. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsListByEnvironment.json
     */
    /**
     * Sample code: List cluster versions by environment.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void listClusterVersionsByEnvironment(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager
            .clusterVersions()
            .listByEnvironmentWithResponse(
                "eastus", ClusterVersionsEnvironment.WINDOWS, com.azure.core.util.Context.NONE);
    }
}
