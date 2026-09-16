
/**
 * Samples for IncidentTasks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/IncidentTasks/IncidentTasks_List.json
     */
    /**
     * Sample code: IncidentTasks_List.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void incidentTasksList(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidentTasks().list("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}
