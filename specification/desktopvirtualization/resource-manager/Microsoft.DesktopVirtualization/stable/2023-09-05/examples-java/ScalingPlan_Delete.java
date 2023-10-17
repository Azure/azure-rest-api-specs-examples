/** Samples for ScalingPlans Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlan_Delete.json
     */
    /**
     * Sample code: ScalingPlans_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlansDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .scalingPlans()
            .deleteByResourceGroupWithResponse("resourceGroup1", "scalingPlan1", com.azure.core.util.Context.NONE);
    }
}
