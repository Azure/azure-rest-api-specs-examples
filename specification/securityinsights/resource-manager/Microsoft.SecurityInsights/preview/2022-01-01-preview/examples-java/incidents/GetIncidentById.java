import com.azure.core.util.Context;

/** Samples for Incidents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/incidents/GetIncidentById.json
     */
    /**
     * Sample code: Get an incident.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnIncident(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidents()
            .getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", Context.NONE);
    }
}
