
/**
 * Samples for ConfigTemplates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ConfigTemplates_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ConfigTemplates_Delete_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void configTemplatesDeleteMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.configTemplates().delete("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
