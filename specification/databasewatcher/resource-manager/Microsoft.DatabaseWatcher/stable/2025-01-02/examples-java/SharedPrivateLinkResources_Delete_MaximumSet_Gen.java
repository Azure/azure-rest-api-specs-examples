
/**
 * Samples for SharedPrivateLinkResources Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/SharedPrivateLinkResources_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: SharedPrivateLinkResources_Delete_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void sharedPrivateLinkResourcesDeleteMaximumSet(
        com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.sharedPrivateLinkResources().delete("apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed",
            com.azure.core.util.Context.NONE);
    }
}
