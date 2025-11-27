
/**
 * Samples for ManagedClusters ListFaultSimulation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/faultSimulation/ManagedClusterListFaultSimulation_example.json
     */
    /**
     * Sample code: List Managed Cluster Fault Simulation.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listManagedClusterFaultSimulation(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusters().listFaultSimulation("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
