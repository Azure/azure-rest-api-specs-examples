/** Samples for ScalingPlanPersonalSchedules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_List.json
     */
    /**
     * Sample code: ScalingPlanPersonalSchedules_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlanPersonalSchedulesList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .scalingPlanPersonalSchedules()
            .list("resourceGroup1", "scalingPlan", 10, true, 0, com.azure.core.util.Context.NONE);
    }
}
