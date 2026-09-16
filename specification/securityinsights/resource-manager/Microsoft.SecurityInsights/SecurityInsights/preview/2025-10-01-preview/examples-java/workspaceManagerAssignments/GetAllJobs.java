
/**
 * Samples for WorkspaceManagerAssignmentJobs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerAssignments/GetAllJobs.json
     */
    /**
     * Sample code: Get all jobs for the specified Sentinel workspace manager assignment.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllJobsForTheSpecifiedSentinelWorkspaceManagerAssignment(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerAssignmentJobs().list("myRg", "myWorkspace", "47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58",
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
