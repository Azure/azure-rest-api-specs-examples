import com.azure.core.util.Context;

/** Samples for Incidents Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/incidents/DeleteIncident.json
     */
    /**
     * Sample code: Delete an incident.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnIncident(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidents()
            .deleteWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", Context.NONE);
    }
}
