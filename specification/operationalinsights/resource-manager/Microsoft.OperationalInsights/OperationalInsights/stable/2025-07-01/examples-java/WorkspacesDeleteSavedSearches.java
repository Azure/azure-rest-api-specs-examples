
/**
 * Samples for SavedSearches Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesDeleteSavedSearches.json
     */
    /**
     * Sample code: SavedSearchesDelete.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void savedSearchesDelete(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.savedSearches().deleteWithResponse("TestRG", "TestWS", "00000000-0000-0000-0000-00000000000",
            com.azure.core.util.Context.NONE);
    }
}
