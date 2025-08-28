
/**
 * Samples for Jobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Jobs_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Jobs_Get_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void
        jobsGetMaximumSet(com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.jobs().getWithResponse("gt", "jobsName", com.azure.core.util.Context.NONE);
    }
}
