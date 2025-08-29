
/**
 * Samples for SolutionTemplateVersions ListBySolutionTemplate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionTemplateVersions_ListBySolutionTemplate_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionTemplateVersions_ListBySolutionTemplate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionTemplateVersionsListBySolutionTemplateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionTemplateVersions().listBySolutionTemplate("rgconfigurationmanager", "testname",
            com.azure.core.util.Context.NONE);
    }
}
