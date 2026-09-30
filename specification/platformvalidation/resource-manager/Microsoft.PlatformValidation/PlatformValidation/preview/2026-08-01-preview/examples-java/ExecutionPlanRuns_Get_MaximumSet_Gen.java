
/**
 * Samples for ExecutionPlanRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ExecutionPlanRuns_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExecutionPlanRuns_Get_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void
        executionPlanRunsGetMaximumSet(com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.executionPlanRuns().getWithResponse("rgvalidate", "cvtest01", "contoso-linux-cert", "run-001",
            com.azure.core.util.Context.NONE);
    }
}
