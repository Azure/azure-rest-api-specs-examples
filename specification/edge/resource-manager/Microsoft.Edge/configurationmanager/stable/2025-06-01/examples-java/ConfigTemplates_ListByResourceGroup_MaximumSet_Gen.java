
/**
 * Samples for ConfigTemplates ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ConfigTemplates_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: ConfigTemplates_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void configTemplatesListByResourceGroupMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.configTemplates().listByResourceGroup("rgconfigurationmanager", com.azure.core.util.Context.NONE);
    }
}
