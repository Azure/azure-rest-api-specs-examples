
/**
 * Samples for NodeTypes ListFaultSimulation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/faultSimulation/NodeTypeListFaultSimulation_example.json
     */
    /**
     * Sample code: List Node Type Fault Simulation.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listNodeTypeFaultSimulation(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.nodeTypes().listFaultSimulation("resRg", "myCluster", "BE", com.azure.core.util.Context.NONE);
    }
}
