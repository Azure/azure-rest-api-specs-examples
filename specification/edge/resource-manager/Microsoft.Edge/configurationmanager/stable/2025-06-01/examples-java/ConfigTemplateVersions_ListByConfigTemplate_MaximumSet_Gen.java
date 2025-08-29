
/**
 * Samples for ConfigTemplateVersions ListByConfigTemplate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ConfigTemplateVersions_ListByConfigTemplate_MaximumSet_Gen.json
     */
    /**
     * Sample code: ConfigTemplateVersions_ListByConfigTemplate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void configTemplateVersionsListByConfigTemplateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.configTemplateVersions().listByConfigTemplate("rgconfigurationmanager", "testname",
            com.azure.core.util.Context.NONE);
    }
}
