
/**
 * Samples for CloudValidations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/CloudValidations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: CloudValidations_Get_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void
        cloudValidationsGetMaximumSet(com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.cloudValidations().getByResourceGroupWithResponse("rgvalidate", "cvtest01",
            com.azure.core.util.Context.NONE);
    }
}
