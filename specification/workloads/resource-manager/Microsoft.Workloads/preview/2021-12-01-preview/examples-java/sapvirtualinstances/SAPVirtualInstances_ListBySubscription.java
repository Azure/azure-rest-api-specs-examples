import com.azure.core.util.Context;

/** Samples for SapVirtualInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_ListBySubscription.json
     */
    /**
     * Sample code: SAPVirtualInstances_ListBySubscription.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesListBySubscription(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapVirtualInstances().list(Context.NONE);
    }
}
