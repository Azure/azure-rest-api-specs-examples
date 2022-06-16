/** Samples for IncidentComments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/incidents/comments/CreateIncidentComment.json
     */
    /**
     * Sample code: Creates or updates an incident comment.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnIncidentComment(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidentComments()
            .define("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014")
            .withExistingIncident("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5")
            .withMessage("Some message")
            .create();
    }
}
