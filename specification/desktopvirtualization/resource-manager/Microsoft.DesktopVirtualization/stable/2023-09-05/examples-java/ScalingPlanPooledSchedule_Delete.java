/** Samples for ScalingPlanPooledSchedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPooledSchedule_Delete.json
     */
    /**
     * Sample code: ScalingPlanPooledSchedules_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlanPooledSchedulesDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .scalingPlanPooledSchedules()
            .deleteWithResponse(
                "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", com.azure.core.util.Context.NONE);
    }
}
