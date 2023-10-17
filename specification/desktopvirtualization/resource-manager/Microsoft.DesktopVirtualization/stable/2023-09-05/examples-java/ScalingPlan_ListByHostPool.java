/** Samples for ScalingPlans ListByHostPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlan_ListByHostPool.json
     */
    /**
     * Sample code: ScalingPlan_ListByHostPool.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlanListByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .scalingPlans()
            .listByHostPool("resourceGroup1", "hostPool1", 10, true, 0, com.azure.core.util.Context.NONE);
    }
}
