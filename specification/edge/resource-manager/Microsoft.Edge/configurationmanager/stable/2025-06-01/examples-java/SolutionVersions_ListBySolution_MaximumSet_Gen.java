
/**
 * Samples for SolutionVersions ListBySolution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionVersions_ListBySolution_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionVersions_ListBySolution_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionVersionsListBySolutionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionVersions().listBySolution("rgconfigurationmanager", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
