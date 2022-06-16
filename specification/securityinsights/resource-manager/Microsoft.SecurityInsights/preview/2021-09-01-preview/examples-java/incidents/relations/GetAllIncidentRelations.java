import com.azure.core.util.Context;

/** Samples for IncidentRelations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/incidents/relations/GetAllIncidentRelations.json
     */
    /**
     * Sample code: Get all incident relations.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllIncidentRelations(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidentRelations()
            .list("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", null, null, null, null, Context.NONE);
    }
}
