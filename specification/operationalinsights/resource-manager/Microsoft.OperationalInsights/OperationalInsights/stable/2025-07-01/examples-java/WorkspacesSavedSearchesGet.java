
/**
 * Samples for SavedSearches Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesSavedSearchesGet.json
     */
    /**
     * Sample code: SavedSearchesGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void savedSearchesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.savedSearches().getWithResponse("TestRG", "TestWS", "00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
