
/**
 * Samples for Incidents ListEntities.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * incidents/entities/GetAllIncidentEntities.json
     */
    /**
     * Sample code: Gets all incident related entities.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getsAllIncidentRelatedEntities(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().listEntitiesWithResponse("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812",
            com.azure.core.util.Context.NONE);
    }
}
