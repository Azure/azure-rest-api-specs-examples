
/**
 * Samples for Incidents ListBookmarks.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/IncidentBookmarks/Incidents_ListBookmarks.json
     */
    /**
     * Sample code: Incidents_ListBookmarks.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        incidentsListBookmarks(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().listBookmarksWithResponse("myRg", "myWorkspace", "69a30280-6a4c-4aa7-9af0-5d63f335d600",
            com.azure.core.util.Context.NONE);
    }
}
