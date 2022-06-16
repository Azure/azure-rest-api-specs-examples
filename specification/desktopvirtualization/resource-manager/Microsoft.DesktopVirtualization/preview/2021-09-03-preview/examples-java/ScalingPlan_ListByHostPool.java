import com.azure.core.util.Context;

/** Samples for ScalingPlans ListByHostPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ScalingPlan_ListByHostPool.json
     */
    /**
     * Sample code: ScalingPlan_ListByHostPool.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlanListByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.scalingPlans().listByHostPool("resourceGroup1", "hostPool1", Context.NONE);
    }
}
