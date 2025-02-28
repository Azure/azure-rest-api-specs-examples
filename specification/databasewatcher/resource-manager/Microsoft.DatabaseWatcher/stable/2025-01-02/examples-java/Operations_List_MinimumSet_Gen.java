
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet.
     * 
     * @param manager Entry point to DatabaseWatcherManager.
     */
    public static void
        operationsListMinimumSet(com.azure.resourcemanager.databasewatcher.DatabaseWatcherManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
