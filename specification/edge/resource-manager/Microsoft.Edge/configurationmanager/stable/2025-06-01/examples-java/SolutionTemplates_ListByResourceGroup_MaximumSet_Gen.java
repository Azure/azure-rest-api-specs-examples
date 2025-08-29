
/**
 * Samples for SolutionTemplates ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionTemplates_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionTemplates_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionTemplatesListByResourceGroupMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionTemplates().listByResourceGroup("rgconfigurationmanager", com.azure.core.util.Context.NONE);
    }
}
