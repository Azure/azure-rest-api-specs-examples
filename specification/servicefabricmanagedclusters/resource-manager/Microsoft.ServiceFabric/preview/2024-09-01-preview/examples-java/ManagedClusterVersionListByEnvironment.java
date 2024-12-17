
import com.azure.resourcemanager.servicefabricmanagedclusters.models.ManagedClusterVersionEnvironment;

/**
 * Samples for ManagedClusterVersion ListByEnvironment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/
     * examples/ManagedClusterVersionListByEnvironment.json
     */
    /**
     * Sample code: List cluster versions by environment.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listClusterVersionsByEnvironment(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusterVersions().listByEnvironmentWithResponse("eastus",
            ManagedClusterVersionEnvironment.WINDOWS, com.azure.core.util.Context.NONE);
    }
}
