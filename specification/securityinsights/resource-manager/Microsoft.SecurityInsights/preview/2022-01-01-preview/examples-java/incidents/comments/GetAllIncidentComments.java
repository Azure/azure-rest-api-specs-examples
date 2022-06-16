import com.azure.core.util.Context;

/** Samples for IncidentComments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/incidents/comments/GetAllIncidentComments.json
     */
    /**
     * Sample code: Get all incident comments.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllIncidentComments(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidentComments()
            .list("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", null, null, null, null, Context.NONE);
    }
}
