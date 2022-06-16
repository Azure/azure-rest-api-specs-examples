import com.azure.core.util.Context;

/** Samples for ScalingPlans Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ScalingPlan_Delete.json
     */
    /**
     * Sample code: ScalingPlans_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlansDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.scalingPlans().deleteWithResponse("resourceGroup1", "scalingPlan1", Context.NONE);
    }
}
