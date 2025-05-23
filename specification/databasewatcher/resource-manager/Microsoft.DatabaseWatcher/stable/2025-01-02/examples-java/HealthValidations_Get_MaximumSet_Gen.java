
/**
 * Samples for HealthValidations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/HealthValidations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: HealthValidations_Get_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        healthValidationsGetMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.healthValidations().getWithResponse("rgWatcher", "testWatcher", "testHealthValidation",
            com.azure.core.util.Context.NONE);
    }
}
