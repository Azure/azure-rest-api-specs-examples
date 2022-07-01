import com.azure.core.util.Context;

/** Samples for Monitors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/monitors_List.json
     */
    /**
     * Sample code: List all SAP monitors in a subscription.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void listAllSAPMonitorsInASubscription(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.monitors().list(Context.NONE);
    }
}
