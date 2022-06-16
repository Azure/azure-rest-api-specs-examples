import com.azure.core.util.Context;

/** Samples for Incidents List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/incidents/GetIncidents.json
     */
    /**
     * Sample code: Get all incidents.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllIncidents(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().list("myRg", "myWorkspace", null, "properties/createdTimeUtc desc", 1, null, Context.NONE);
    }
}
