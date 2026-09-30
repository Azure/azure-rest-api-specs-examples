
/**
 * Samples for ExecutionPlanRuns ListByExecutionPlan.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ExecutionPlanRuns_ListByExecutionPlan_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExecutionPlanRuns_ListByExecutionPlan_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void executionPlanRunsListByExecutionPlanMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.executionPlanRuns().listByExecutionPlan("rgvalidate", "cvtest01", "contoso-linux-cert", null,
            com.azure.core.util.Context.NONE);
    }
}
