
/**
 * Samples for Watchers Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Watchers_Start_MaximumSet_Gen.json
     */
    /**
     * Sample code: Watchers_Start_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        watchersStartMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.watchers().start("rgWatcher", "testWatcher", com.azure.core.util.Context.NONE);
    }
}
