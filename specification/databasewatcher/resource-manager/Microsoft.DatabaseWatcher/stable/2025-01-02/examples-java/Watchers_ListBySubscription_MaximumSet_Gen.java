
/**
 * Samples for Watchers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Watchers_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Watchers_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        watchersListBySubscriptionMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.watchers().list(com.azure.core.util.Context.NONE);
    }
}
