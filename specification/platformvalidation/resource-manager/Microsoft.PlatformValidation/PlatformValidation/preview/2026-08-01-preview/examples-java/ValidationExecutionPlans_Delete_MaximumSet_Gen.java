
/**
 * Samples for ValidationExecutionPlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationExecutionPlans_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ValidationExecutionPlans_Delete_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationExecutionPlansDeleteMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationExecutionPlans().delete("rgvalidate", "cvtest01", "veptest01",
            com.azure.core.util.Context.NONE);
    }
}
