
/**
 * Samples for GoalAssignments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalAssignments_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: GoalAssignments_Get_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalAssignmentsGetMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalAssignments().getWithResponse("production-sg", "zonal-resiliency-goal",
            com.azure.core.util.Context.NONE);
    }
}
