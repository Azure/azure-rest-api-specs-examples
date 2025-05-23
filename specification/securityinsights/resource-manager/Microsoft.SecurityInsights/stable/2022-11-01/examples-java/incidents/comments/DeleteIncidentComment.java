
/**
 * Samples for IncidentComments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/incidents/
     * comments/DeleteIncidentComment.json
     */
    /**
     * Sample code: Delete the incident comment.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteTheIncidentComment(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidentComments().deleteWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014", com.azure.core.util.Context.NONE);
    }
}
