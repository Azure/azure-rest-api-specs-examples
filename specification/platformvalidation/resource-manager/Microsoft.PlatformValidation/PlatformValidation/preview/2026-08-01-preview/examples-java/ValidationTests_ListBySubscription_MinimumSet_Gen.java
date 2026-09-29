
/**
 * Samples for ValidationTests List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ValidationTests_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: ValidationTests_ListBySubscription_MinimumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void validationTestsListBySubscriptionMinimumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.validationTests().list(null, com.azure.core.util.Context.NONE);
    }
}
