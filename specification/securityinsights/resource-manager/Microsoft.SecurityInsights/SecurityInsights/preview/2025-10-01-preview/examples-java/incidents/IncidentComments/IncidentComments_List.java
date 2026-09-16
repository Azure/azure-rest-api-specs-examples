
/**
 * Samples for IncidentComments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/IncidentComments/IncidentComments_List.json
     */
    /**
     * Sample code: IncidentComments_List.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        incidentCommentsList(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidentComments().list("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", null, null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
