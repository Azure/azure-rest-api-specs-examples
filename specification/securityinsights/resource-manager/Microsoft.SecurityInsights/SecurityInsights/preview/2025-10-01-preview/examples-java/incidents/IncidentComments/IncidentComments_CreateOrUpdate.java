
/**
 * Samples for IncidentComments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/IncidentComments/IncidentComments_CreateOrUpdate.json
     */
    /**
     * Sample code: IncidentComments_CreateOrUpdate.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        incidentCommentsCreateOrUpdate(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidentComments().define("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014")
            .withExistingIncident("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5")
            .withMessage("Some message").create();
    }
}
