
/**
 * Samples for GoalResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalResources_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: GoalResources_Get_MinimumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        goalResourcesGetMinimumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalResources().getWithResponse("production-sg", "zonal-resiliency-goal", "primary-vm",
            com.azure.core.util.Context.NONE);
    }
}
