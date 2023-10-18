/** Samples for ScalingPlanPersonalSchedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_Delete.json
     */
    /**
     * Sample code: ScalingPlanPersonalSchedules_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlanPersonalSchedulesDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .scalingPlanPersonalSchedules()
            .deleteWithResponse(
                "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", com.azure.core.util.Context.NONE);
    }
}
