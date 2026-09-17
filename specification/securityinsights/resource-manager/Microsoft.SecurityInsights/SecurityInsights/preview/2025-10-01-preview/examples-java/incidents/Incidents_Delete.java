
/**
 * Samples for Incidents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/Incidents_Delete.json
     */
    /**
     * Sample code: Incidents_Delete.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void incidentsDelete(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().deleteWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}
