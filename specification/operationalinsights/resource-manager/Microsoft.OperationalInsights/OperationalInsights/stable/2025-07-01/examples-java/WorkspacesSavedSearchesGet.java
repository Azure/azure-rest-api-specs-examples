
/**
 * Samples for SavedSearches Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesSavedSearchesGet.json
     */
    /**
     * Sample code: SavedSearchesGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void savedSearchesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.savedSearches().getWithResponse("TestRG", "TestWS", "00000000-0000-0000-0000-00000000000",
            com.azure.core.util.Context.NONE);
    }
}
