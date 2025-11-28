
import com.azure.resourcemanager.servicefabricmanagedclusters.models.FaultSimulationIdContent;

/**
 * Samples for ManagedClusters GetFaultSimulation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/faultSimulation/ManagedClusterGetFaultSimulation_example.json
     */
    /**
     * Sample code: Get Managed Cluster Fault Simulation.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getManagedClusterFaultSimulation(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusters().getFaultSimulationWithResponse("resRg", "myCluster",
            new FaultSimulationIdContent().withSimulationId("aec13cc2-1d39-4ba6-a1a8-2fc35b00643c"),
            com.azure.core.util.Context.NONE);
    }
}
