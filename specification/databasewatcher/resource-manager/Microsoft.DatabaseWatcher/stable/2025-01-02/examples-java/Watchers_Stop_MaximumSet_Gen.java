
/**
 * Samples for Watchers Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Watchers_Stop_MaximumSet_Gen.json
     */
    /**
     * Sample code: Watchers_Stop_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        watchersStopMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.watchers().stop("rgWatcher", "myWatcher", com.azure.core.util.Context.NONE);
    }
}
