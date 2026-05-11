
/**
 * Samples for SavedSearches Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesDeleteSavedSearches.json
     */
    /**
     * Sample code: SavedSearchesDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void savedSearchesDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.savedSearches().deleteWithResponse("TestRG", "TestWS", "00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
