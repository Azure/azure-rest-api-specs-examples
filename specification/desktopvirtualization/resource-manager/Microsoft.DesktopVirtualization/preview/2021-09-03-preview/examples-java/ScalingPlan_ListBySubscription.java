import com.azure.core.util.Context;

/** Samples for ScalingPlans List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ScalingPlan_ListBySubscription.json
     */
    /**
     * Sample code: ScalingPlans_ListBySubscription.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void scalingPlansListBySubscription(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.scalingPlans().list(Context.NONE);
    }
}
