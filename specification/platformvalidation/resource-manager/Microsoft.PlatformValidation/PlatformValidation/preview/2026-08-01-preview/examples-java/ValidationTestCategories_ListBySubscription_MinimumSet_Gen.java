
/**
 * Samples for ValidationTestCategories List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationTestCategories_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: ValidationTestCategories_ListBySubscription_MinimumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationTestCategoriesListBySubscriptionMinimumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationTestCategories().list(null, com.azure.core.util.Context.NONE);
    }
}
