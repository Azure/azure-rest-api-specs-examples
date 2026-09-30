
/**
 * Samples for CloudValidations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/CloudValidations_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: CloudValidations_ListBySubscription_MinimumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void cloudValidationsListBySubscriptionMinimumSet(
        com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.cloudValidations().list(null, com.azure.core.util.Context.NONE);
    }
}
