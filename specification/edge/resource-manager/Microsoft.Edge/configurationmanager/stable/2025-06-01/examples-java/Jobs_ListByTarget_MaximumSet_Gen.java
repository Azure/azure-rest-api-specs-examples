
/**
 * Samples for Jobs ListByTarget.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Jobs_ListByTarget_MaximumSet_Gen.json
     */
    /**
     * Sample code: Jobs_ListByTarget_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void jobsListByTargetMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.jobs().listByTarget("gt", com.azure.core.util.Context.NONE);
    }
}
