
/**
 * Samples for IncidentTasks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/IncidentTasks/IncidentTasks_Get.json
     */
    /**
     * Sample code: IncidentTasks_Get.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void incidentTasksGet(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidentTasks().getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014", com.azure.core.util.Context.NONE);
    }
}
