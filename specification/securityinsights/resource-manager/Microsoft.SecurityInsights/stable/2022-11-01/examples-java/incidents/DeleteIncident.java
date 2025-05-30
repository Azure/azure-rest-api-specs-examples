
/**
 * Samples for Incidents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/incidents/
     * DeleteIncident.json
     */
    /**
     * Sample code: Delete an incident.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnIncident(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().deleteWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}
