
/**
 * Samples for ValidationTestCategories List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationTestCategories_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: ValidationTestCategories_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationTestCategoriesListBySubscriptionMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationTestCategories().list("audience eq 'Public'", com.azure.core.util.Context.NONE);
    }
}
