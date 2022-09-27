import com.azure.core.util.Context;

/** Samples for Incidents ListAlerts. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/incidents/GetAllIncidentAlerts.json
     */
    /**
     * Sample code: Get all incident alerts.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllIncidentAlerts(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidents()
            .listAlertsWithResponse("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", Context.NONE);
    }
}
