
/**
 * Samples for GoalAssignments RefreshGoalResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalAssignments_RefreshGoalResources_MaximumSet_Gen.json
     */
    /**
     * Sample code: GoalAssignments_RefreshGoalResources_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalAssignmentsRefreshGoalResourcesMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalAssignments().refreshGoalResources("production-sg", "zonal-resiliency-goal",
            com.azure.core.util.Context.NONE);
    }
}
