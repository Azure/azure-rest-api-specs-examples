
/**
 * Samples for Watchers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Watchers_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Watchers_Delete_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        watchersDeleteMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.watchers().delete("rgWatcher", "testWatcher", com.azure.core.util.Context.NONE);
    }
}
