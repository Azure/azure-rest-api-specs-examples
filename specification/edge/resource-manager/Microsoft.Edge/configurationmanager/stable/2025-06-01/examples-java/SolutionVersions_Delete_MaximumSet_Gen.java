
/**
 * Samples for SolutionVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionVersions_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionVersions_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionVersionsDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionVersions().delete("rgconfigurationmanager", "testname", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
