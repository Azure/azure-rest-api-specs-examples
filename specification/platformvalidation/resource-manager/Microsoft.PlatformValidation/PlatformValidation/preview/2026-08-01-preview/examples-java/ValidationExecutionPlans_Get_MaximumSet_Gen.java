
/**
 * Samples for ValidationExecutionPlans Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationExecutionPlans_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ValidationExecutionPlans_Get_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationExecutionPlansGetMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationExecutionPlans().getWithResponse("rgvalidate", "cvtest01", "contoso-linux-cert",
            com.azure.core.util.Context.NONE);
    }
}
