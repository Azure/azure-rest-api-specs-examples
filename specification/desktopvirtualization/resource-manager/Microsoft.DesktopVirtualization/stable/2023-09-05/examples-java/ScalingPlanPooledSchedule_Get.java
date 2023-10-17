/** Samples for ScalingPlanPooledSchedules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPooledSchedule_Get.json
     */
    /**
     * Sample code: ScalingPlanPooledSchedules_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlanPooledSchedulesGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .scalingPlanPooledSchedules()
            .getWithResponse(
                "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", com.azure.core.util.Context.NONE);
    }
}
