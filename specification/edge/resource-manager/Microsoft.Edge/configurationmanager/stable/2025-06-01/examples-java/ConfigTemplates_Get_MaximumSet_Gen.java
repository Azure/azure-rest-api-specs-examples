
/**
 * Samples for ConfigTemplates GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ConfigTemplates_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ConfigTemplates_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void configTemplatesGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.configTemplates().getByResourceGroupWithResponse("rgconfigurationmanager", "testname",
            com.azure.core.util.Context.NONE);
    }
}
