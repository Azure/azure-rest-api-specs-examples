
/**
 * Samples for Instances ListBySolution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Instances_ListBySolution_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instances_ListBySolution_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void instancesListBySolutionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.instances().listBySolution("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
