
/**
 * Samples for Incidents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/Incidents_List.json
     */
    /**
     * Sample code: Incidents_List.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void incidentsList(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().list("myRg", "myWorkspace", null, "properties/createdTimeUtc desc", 1, null,
            com.azure.core.util.Context.NONE);
    }
}
