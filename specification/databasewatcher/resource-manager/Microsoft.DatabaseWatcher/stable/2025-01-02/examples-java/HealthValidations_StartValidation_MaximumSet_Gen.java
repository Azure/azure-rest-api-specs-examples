
/**
 * Samples for HealthValidations StartValidation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/HealthValidations_StartValidation_MaximumSet_Gen.json
     */
    /**
     * Sample code: HealthValidations_StartValidation_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void healthValidationsStartValidationMaximumSet(
        com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.healthValidations().startValidation("rgWatcher", "testWatcher", "testHealthValidation",
            com.azure.core.util.Context.NONE);
    }
}
