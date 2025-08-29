
/**
 * Samples for SolutionTemplateVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionTemplateVersions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionTemplateVersions_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionTemplateVersionsGetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionTemplateVersions().getWithResponse("rgconfigurationmanager", "testname", "1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
