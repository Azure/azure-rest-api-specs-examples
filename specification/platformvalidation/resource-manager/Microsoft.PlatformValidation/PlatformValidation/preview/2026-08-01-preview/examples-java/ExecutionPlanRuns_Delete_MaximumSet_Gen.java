
/**
 * Samples for ExecutionPlanRuns Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ExecutionPlanRuns_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExecutionPlanRuns_Delete_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void executionPlanRunsDeleteMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.executionPlanRuns().delete("rgvalidate", "cvtest01", "veptest01", "veprun01",
            com.azure.core.util.Context.NONE);
    }
}
