
/**
 * Samples for SolutionVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionVersions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionVersions_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionVersionsGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionVersions().getWithResponse("rgconfigurationmanager", "testname", "testname", "testname",
            com.azure.core.util.Context.NONE);
    }
}
