/** Samples for SavedSearches ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/SavedSearchesListByWorkspace.json
     */
    /**
     * Sample code: SavedSearchesList.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void savedSearchesList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.savedSearches().listByWorkspaceWithResponse("TestRG", "TestWS", com.azure.core.util.Context.NONE);
    }
}
