
/**
 * Samples for ValidationTestRuns Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationTestRuns_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ValidationTestRuns_Get_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationTestRunsGetMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationTestRuns().getWithResponse("rgvalidate", "cvtest01", "contoso-linux-cert", "run-001",
            "linux-quality-run", com.azure.core.util.Context.NONE);
    }
}
