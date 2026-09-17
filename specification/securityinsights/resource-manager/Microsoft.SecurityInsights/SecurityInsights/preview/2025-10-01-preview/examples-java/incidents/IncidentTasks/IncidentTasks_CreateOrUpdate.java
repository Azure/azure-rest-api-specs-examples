
import com.azure.resourcemanager.securityinsights.models.IncidentTaskStatus;

/**
 * Samples for IncidentTasks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/IncidentTasks/IncidentTasks_CreateOrUpdate.json
     */
    /**
     * Sample code: IncidentTasks_CreateOrUpdate.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        incidentTasksCreateOrUpdate(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidentTasks().define("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014")
            .withExistingIncident("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5").withTitle("Task title")
            .withStatus(IncidentTaskStatus.NEW).withDescription("Task description").create();
    }
}
