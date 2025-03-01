
/**
 * Samples for AlertRuleResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/AlertRuleResources_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AlertRuleResources_Get_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        alertRuleResourcesGetMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.alertRuleResources().getWithResponse("rgWatcher", "testWatcher", "testAlert",
            com.azure.core.util.Context.NONE);
    }
}
