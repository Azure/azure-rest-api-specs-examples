/** Samples for ScalingPlanPersonalSchedules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_Get.json
     */
    /**
     * Sample code: ScalingPlanPersonalSchedules_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlanPersonalSchedulesGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .scalingPlanPersonalSchedules()
            .getWithResponse(
                "resourceGroup1",
                "PersonalScalingPlan1",
                "PersonalScalingPlanSchedule",
                com.azure.core.util.Context.NONE);
    }
}
