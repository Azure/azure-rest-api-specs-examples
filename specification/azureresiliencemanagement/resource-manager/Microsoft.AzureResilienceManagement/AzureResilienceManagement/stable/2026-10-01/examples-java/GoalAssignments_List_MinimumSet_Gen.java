
/**
 * Samples for GoalAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalAssignments_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: GoalAssignments_List_MinimumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalAssignmentsListMinimumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalAssignments().list("production-sg", null, null, com.azure.core.util.Context.NONE);
    }
}
