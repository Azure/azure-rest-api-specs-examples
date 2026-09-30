
/**
 * Samples for ValidationExecutionPlans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationExecutionPlans_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: ValidationExecutionPlans_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationExecutionPlansListByResourceGroupMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationExecutionPlans().listByResourceGroup("rgvalidate", "cvtest01", null,
            com.azure.core.util.Context.NONE);
    }
}
