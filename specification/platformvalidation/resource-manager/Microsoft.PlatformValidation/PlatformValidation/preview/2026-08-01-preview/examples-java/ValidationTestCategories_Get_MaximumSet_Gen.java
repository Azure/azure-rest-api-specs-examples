
/**
 * Samples for ValidationTestCategories Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationTestCategories_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ValidationTestCategories_Get_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationTestCategoriesGetMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationTestCategories().getWithResponse("linux-quality-validations",
            com.azure.core.util.Context.NONE);
    }
}
