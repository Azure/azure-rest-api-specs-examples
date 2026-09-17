
/**
 * Samples for IncidentRelations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/relations/GetAllIncidentRelations.json
     */
    /**
     * Sample code: Get all incident relations.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllIncidentRelations(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidentRelations().list("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", null, null,
            null, null, com.azure.core.util.Context.NONE);
    }
}
