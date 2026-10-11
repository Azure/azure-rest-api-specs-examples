
/**
 * Samples for GoalResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/GoalResources_Get_Complete_Example.json
     */
    /**
     * Sample code: GoalResources_Get_Complete_Example.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void goalResourcesGetCompleteExample(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.goalResources().getWithResponse("production-sg", "zonal-resiliency-goal", "primary-vm",
            com.azure.core.util.Context.NONE);
    }
}
