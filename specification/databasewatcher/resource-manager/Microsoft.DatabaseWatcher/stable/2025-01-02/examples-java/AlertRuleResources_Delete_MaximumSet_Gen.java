
/**
 * Samples for AlertRuleResources Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/AlertRuleResources_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AlertRuleResources_Delete_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        alertRuleResourcesDeleteMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.alertRuleResources().deleteWithResponse("rgWatcher", "testWatcher", "testAlert",
            com.azure.core.util.Context.NONE);
    }
}
