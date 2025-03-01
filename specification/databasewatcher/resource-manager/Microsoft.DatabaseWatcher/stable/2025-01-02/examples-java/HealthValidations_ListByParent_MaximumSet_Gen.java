
/**
 * Samples for HealthValidations ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/HealthValidations_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: HealthValidations_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void healthValidationsListByParentMaximumSet(
        com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.healthValidations().listByParent("rgWatcher", "testWatcher", com.azure.core.util.Context.NONE);
    }
}
