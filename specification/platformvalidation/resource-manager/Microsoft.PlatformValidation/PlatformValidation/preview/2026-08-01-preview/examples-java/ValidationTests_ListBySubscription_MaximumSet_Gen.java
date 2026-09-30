
/**
 * Samples for ValidationTests List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationTests_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: ValidationTests_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationTestsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationTests().list("audience eq 'Public'", com.azure.core.util.Context.NONE);
    }
}
