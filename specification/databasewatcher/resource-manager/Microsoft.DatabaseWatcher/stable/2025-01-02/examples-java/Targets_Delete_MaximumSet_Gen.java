
/**
 * Samples for Targets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Targets_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_Delete_MaximumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        targetsDeleteMaximumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.targets().deleteWithResponse("apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed",
            com.azure.core.util.Context.NONE);
    }
}
