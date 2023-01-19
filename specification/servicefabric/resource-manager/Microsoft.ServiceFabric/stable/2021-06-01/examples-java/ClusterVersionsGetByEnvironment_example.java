import com.azure.resourcemanager.servicefabric.models.ClusterVersionsEnvironment;

/** Samples for ClusterVersions GetByEnvironment. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsGetByEnvironment_example.json
     */
    /**
     * Sample code: Get cluster version by environment.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getClusterVersionByEnvironment(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager
            .clusterVersions()
            .getByEnvironmentWithResponse(
                "eastus", ClusterVersionsEnvironment.WINDOWS, "6.1.480.9494", com.azure.core.util.Context.NONE);
    }
}
