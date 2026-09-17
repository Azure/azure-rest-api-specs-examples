
/**
 * Samples for WorkspaceManagerAssignmentJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerAssignments/GetJob.json
     */
    /**
     * Sample code: Get a workspace manager job.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAWorkspaceManagerJob(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerAssignmentJobs().getWithResponse("myRg", "myWorkspace",
            "47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58", "cfbe1338-8276-4d5d-8b96-931117f9fa0e",
            com.azure.core.util.Context.NONE);
    }
}
