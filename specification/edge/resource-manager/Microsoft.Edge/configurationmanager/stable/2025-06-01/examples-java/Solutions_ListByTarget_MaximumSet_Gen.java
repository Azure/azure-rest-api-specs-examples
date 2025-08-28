
/**
 * Samples for Solutions ListByTarget.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Solutions_ListByTarget_MaximumSet_Gen.json
     */
    /**
     * Sample code: Solutions_ListByTarget_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionsListByTargetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutions().listByTarget("rgconfigurationmanager", "testname", com.azure.core.util.Context.NONE);
    }
}
