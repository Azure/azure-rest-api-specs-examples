
/**
 * Samples for SolutionTemplates GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionTemplates_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionTemplates_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionTemplatesGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionTemplates().getByResourceGroupWithResponse("rgconfigurationmanager", "testname",
            com.azure.core.util.Context.NONE);
    }
}
