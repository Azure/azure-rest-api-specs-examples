
/**
 * Samples for GoalResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalResources_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: GoalResources_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalResourcesListMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalResources().list("production-sg", "zonal-resiliency-goal", "xntbyoswztnmvitj", 69,
            com.azure.core.util.Context.NONE);
    }
}
