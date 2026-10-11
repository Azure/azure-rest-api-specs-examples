
/**
 * Samples for GoalAssignments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalAssignments_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: GoalAssignments_Delete_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalAssignmentsDeleteMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalAssignments().delete("production-sg", "zonal-resiliency-goal", com.azure.core.util.Context.NONE);
    }
}
