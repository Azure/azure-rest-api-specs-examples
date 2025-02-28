
/**
 * Samples for AlertRuleResources ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/AlertRuleResources_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: AlertRuleResources_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void alertRuleResourcesListByParentMaximumSet(
        com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.alertRuleResources().listByParent("rgWatcher", "testWatcher", com.azure.core.util.Context.NONE);
    }
}
