
/**
 * Samples for SolutionTemplates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionTemplates_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionTemplates_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionTemplatesDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionTemplates().delete("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
