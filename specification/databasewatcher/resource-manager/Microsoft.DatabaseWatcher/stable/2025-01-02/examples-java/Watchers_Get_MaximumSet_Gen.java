
/**
 * Samples for Watchers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Watchers_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Watchers_Get_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void watchersGetMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.watchers().getByResourceGroupWithResponse("rgWatcher", "myWatcher", com.azure.core.util.Context.NONE);
    }
}
