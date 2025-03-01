
/**
 * Samples for SharedPrivateLinkResources ListByWatcher.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/SharedPrivateLinkResources_ListByWatcher_MaximumSet_Gen.json
     */
    /**
     * Sample code: SharedPrivateLinkResources_ListByWatcher_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void sharedPrivateLinkResourcesListByWatcherMaximumSet(
        com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.sharedPrivateLinkResources().listByWatcher("apiTest-ddat4p", "databasemo3ej9ih",
            com.azure.core.util.Context.NONE);
    }
}
