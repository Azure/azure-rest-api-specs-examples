
/**
 * Samples for Targets ListByWatcher.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Targets_ListByWatcher_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_ListByWatcher_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        targetsListByWatcherMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.targets().listByWatcher("apiTest-ddat4p", "databasemo3ej9ih", com.azure.core.util.Context.NONE);
    }
}
