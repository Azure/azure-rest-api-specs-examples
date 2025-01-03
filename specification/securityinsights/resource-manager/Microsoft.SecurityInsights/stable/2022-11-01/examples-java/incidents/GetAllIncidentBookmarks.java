
/**
 * Samples for Incidents ListBookmarks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/incidents/
     * GetAllIncidentBookmarks.json
     */
    /**
     * Sample code: Get all incident bookmarks.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllIncidentBookmarks(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().listBookmarksWithResponse("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812",
            com.azure.core.util.Context.NONE);
    }
}
