
import com.azure.resourcemanager.servicefabricmanagedclusters.models.ManagedClusterVersionEnvironment;

/**
 * Samples for ManagedClusterVersion ListByEnvironment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ManagedClusterVersionListByEnvironment.json
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
