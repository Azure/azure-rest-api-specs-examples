
/**
 * Samples for GoalAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalAssignments_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: GoalAssignments_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalAssignmentsListMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalAssignments().list("production-sg", "xntbyoswztnmvitj", 69, com.azure.core.util.Context.NONE);
    }
}
