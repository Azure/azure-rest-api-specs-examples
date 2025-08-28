
/**
 * Samples for ConfigTemplateVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ConfigTemplateVersions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ConfigTemplateVersions_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void configTemplateVersionsGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.configTemplateVersions().getWithResponse("rgconfigurationmanager", "testname", "1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
