
import com.azure.resourcemanager.servicefabricmanagedclusters.models.FaultSimulationContentWrapper;
import com.azure.resourcemanager.servicefabricmanagedclusters.models.ZoneFaultSimulationContent;
import java.util.Arrays;

/**
 * Samples for ManagedClusters StartFaultSimulation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/faultSimulation/ManagedClusterStartFaultSimulation_example.json
     */
    /**
     * Sample code: Start Managed Cluster Fault Simulation.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void startManagedClusterFaultSimulation(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedClusters().startFaultSimulation("resRg", "myCluster", new FaultSimulationContentWrapper()
            .withParameters(new ZoneFaultSimulationContent().withZones(Arrays.asList("2"))),
            com.azure.core.util.Context.NONE);
    }
}
